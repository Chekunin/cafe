package memory

import (
	"cafe/pkg/client_sso"
	errs "cafe/pkg/client_sso/errors"
	"cafe/pkg/db_manager"
	dbManagerErrs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type ClientSso struct {
	dbManager db_manager.IDBManager
}

type NewClientSsoParams struct {
	DbManager db_manager.IDBManager
}

var signingKey = "merlinSigningKey"        //TODO cfg
var signingMethod = jwt.SigningMethodHS256 //TODO cfg

type tokenClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func NewClientSso(params NewClientSsoParams) (client_sso.IClientSso, error) {
	clientSso := ClientSso{
		dbManager: params.DbManager,
	}

	return &clientSso, nil
}

func (c ClientSso) Login(ctx context.Context, userName string, password string) (string, error) {
	// todo: здесь же где-то должен делать маппинг ошибок
	user, err := c.dbManager.GetUserByUserName(ctx, userName)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserName username=%s", userName), err)
		if errors.Is(err, dbManagerErrs.ErrorEntityNotFound) {
			err = wrapErr.NewWrapErr(errs.ErrIncorrectLoginOrPassword, err)
		}
		return "", err
	}

	if checkUserPassword(user, password) {
		err := wrapErr.NewWrapErr(errs.ErrIncorrectLoginOrPassword, fmt.Errorf("check password"))
		return "", err
	}

	token, _ := jwt.NewWithClaims(signingMethod, &tokenClaims{
		UserID: user.ID,
	}).SignedString([]byte(signingKey))

	return token, nil
}

func (c ClientSso) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	_, err := c.getTokenClaims(token)
	if err != nil {
		err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
		return false, err
	}
	// todo: тут надо делать проверку на истечение срока действия токена наверное
	return true, nil
}

func (c ClientSso) GetUserID(ctx context.Context, token string) (int, error) {
	claims, err := c.getTokenClaims(token)
	if err != nil {
		err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
		return 0, err
	}
	return claims.UserID, nil
}

func (c ClientSso) getTokenClaims(token string) (*tokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %s", err)
	}

	if signingMethod.Alg() != parsedToken.Header["alg"] {
		return nil, fmt.Errorf("error validating token algorithm: expected %s signing method but token specified %s", signingMethod.Alg(), parsedToken.Header["alg"])
	}

	if claims, ok := parsedToken.Claims.(*tokenClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token is invalid")
}

func checkUserPassword(user models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
