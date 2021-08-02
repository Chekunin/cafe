package usecase

import (
	errs "cafe/pkg/client_gateway_service/errors"
	clientSsoModels "cafe/pkg/client_sso/models"
	dbManagerErrs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func (u *Usecase) Login(ctx context.Context, username string, password string) (clientSsoModels.Tokens, error) {
	tokens, err := u.clientSso.Login(ctx, username, password)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso Login username=%s", username), err)
		return clientSsoModels.Tokens{}, err
	}

	return tokens, nil
}

func (u *Usecase) Logout(ctx context.Context, token string) error {
	if err := u.clientSso.Logout(ctx, token); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso Logout token=%s", token), err)
		return err
	}

	return nil
}

func (u *Usecase) RefreshToken(ctx context.Context, refreshToken string) (clientSsoModels.Tokens, error) {
	tokens, err := u.clientSso.RefreshToken(ctx, refreshToken)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso RefreshToken refreshToken=%s", refreshToken), err)
		return clientSsoModels.Tokens{}, err
	}

	return tokens, nil
}

func (u *Usecase) SignUp(ctx context.Context, username string, phone string, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("bcrypt GenerateFromPassword with DefaultCost"), err)
		return models.User{}, err
	}

	user := models.User{
		Name:     username,
		Phone:    phone,
		Password: string(hashedPassword),
	}

	if _, err := u.dbManager.GetUserByUserName(ctx, username); err == nil {
		err2 := errs.ErrorFieldNotUnique
		err2.Meta = "name"
		err = wrapErr.NewWrapErr(err2, err)
		return models.User{}, err
	} else if !errors.Is(err, dbManagerErrs.ErrorEntityNotFound) {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserName username=%s", username), err)
		return models.User{}, err
	}

	if _, err := u.dbManager.GetUserByVerifiedPhone(ctx, phone); err == nil {
		err2 := errs.ErrorFieldNotUnique
		err2.Meta = "phone"
		err = wrapErr.NewWrapErr(err2, err)
		return models.User{}, err
	} else if !errors.Is(err, dbManagerErrs.ErrorEntityNotFound) {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByVerifiedPhone phone=%s", phone), err)
		return models.User{}, err
	}

	if err := u.dbManager.CreateUser(ctx, &user); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CreateUser=%+v", user), err)
		if errors.Is(err, dbManagerErrs.ErrUniqueViolation(nil)) {
			err = wrapErr.NewWrapErr(errs.ErrorFieldNotUnique, err)
		}
		return models.User{}, err
	}

	approveCode := randRunes(5, []rune("0123456789"))
	userPhoneCode := models.UserPhoneCode{
		UserID:         user.ID,
		Code:           approveCode,
		CreateDatetime: time.Now(),
		Actual:         true,
		LeftAttempts:   4, // todo: вынести в конфиг
	}
	if err := u.dbManager.CreateUserPhoneCode(ctx, &userPhoneCode); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager CreateUserPhoneCode"), err)
		return models.User{}, err
	}

	return user, nil
}

func (u *Usecase) ApprovePhone(ctx context.Context, userID int, phone string, code string) error {
	user, err := u.dbManager.GetUserByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserID userID=%d", userID), err)
		return err
	}

	if user.Phone != phone {
		err := wrapErr.NewWrapErr(fmt.Errorf("user's phone = %s and specified phone = %s", user.Phone, phone), errs.ErrorPhoneNotCorrect)
		return err
	}

	// здесь должны сверять код подтверждения
	userPhoneCode, err := u.dbManager.GetActualUserPhoneCodeByUserID(ctx, userID)
	if err != nil {
		err := wrapErr.NewWrapErr(fmt.Errorf("dbManager GetActualUserPhoneCodeByUserID userID=%d", userID), err)
		return err
	}
	if userPhoneCode.Code != code ||
		userPhoneCode.LeftAttempts <= 0 ||
		!userPhoneCode.Actual ||
		userPhoneCode.CreateDatetime.Add(5*time.Minute).Before(time.Now()) {
		// todo: уменьшать кол-во попыток входа
		err := wrapErr.NewWrapErr(fmt.Errorf("userPhoneCode is incorrect"), errs.ErrorPhoneCodeNotCorrect)
		return err
	}

	// помечаем user_phone_code как не актуальный
	// в users помечаем номер телефона как подтверждённый

	return nil
}

func randRunes(n int, alphabet []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
