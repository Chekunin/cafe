package memory

import (
	"cafe/pkg/client_sso"
	errs "cafe/pkg/client_sso/errors"
	"cafe/pkg/client_sso/memory/auth"
	ssoModels "cafe/pkg/client_sso/models"
	"cafe/pkg/db_manager"
	dbManagerErrs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type ClientSso struct {
	dbManager db_manager.IDBManager
	rd        auth.AuthInterface
	tk        auth.TokenInterface
}

type NewClientSsoParams struct {
	DbManager db_manager.IDBManager
	Redis     *redis.Client
}

func NewClientSso(params NewClientSsoParams) (client_sso.IClientSso, error) {
	rd := auth.NewAuth(params.Redis)
	tk := auth.NewToken()
	clientSso := ClientSso{
		dbManager: params.DbManager,
		rd:        rd,
		tk:        tk,
	}

	return &clientSso, nil
}

func (c ClientSso) Login(ctx context.Context, userName string, password string) (ssoModels.Tokens, error) {
	user, err := c.dbManager.GetUserByUserName(ctx, userName)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager GetUserByUserName username=%s", userName), err)
		if errors.Is(err, dbManagerErrs.ErrorEntityNotFound) {
			err = wrapErr.NewWrapErr(errs.ErrIncorrectLoginOrPassword, err)
		}
		return ssoModels.Tokens{}, err
	}

	if !checkUserPassword(user, password) {
		err := wrapErr.NewWrapErr(errs.ErrIncorrectLoginOrPassword, fmt.Errorf("check password"))
		return ssoModels.Tokens{}, err
	}

	ts, err := c.tk.CreateToken(strconv.Itoa(user.ID))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("tk CreateToken for userID=%d", user.ID), err)
		return ssoModels.Tokens{}, err
	}
	if err := c.rd.CreateAuth(strconv.Itoa(user.ID), ts); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("rs CreateAuth for userID=%d", user.ID), err)
		return ssoModels.Tokens{}, err
	}
	tokens := ssoModels.Tokens{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	return tokens, nil
}

func (c ClientSso) Logout(ctx context.Context, token string) error {
	metadata, err := c.tk.ParseAccessDetails(token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("tk ParseAccessDetails token=%s", token), err)
		return err
	}
	if err := c.rd.DeleteTokens(metadata); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("rd DeleteTokens"), err)
		return err
	}
	return nil
}

func (c ClientSso) RefreshToken(ctx context.Context, refreshToken string) (ssoModels.Tokens, error) {
	token, err := auth.VerifyToken(refreshToken)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("VerifyToken token=%s", refreshToken), err)
		return ssoModels.Tokens{}, err
	}
	claims, ok := token.Claims.(*auth.TokenClaims)
	if !ok || !token.Valid {
		err := wrapErr.NewWrapErr(fmt.Errorf("token is invalid or incorrect"), nil)
		return ssoModels.Tokens{}, err
	}

	if err := c.rd.DeleteRefresh(claims.TokenUuid); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("DeleteRefresh tokenUuid=%s", claims.TokenUuid), err)
		return ssoModels.Tokens{}, err
	}
	tokenDetails, err := c.tk.CreateToken(claims.UserID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("CreateToken userID=%s", claims.UserID), err)
		return ssoModels.Tokens{}, err
	}
	if err := c.rd.CreateAuth(claims.UserID, tokenDetails); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("CreateAuth userID=%s, tokenDetails=%+v", claims.UserID, tokenDetails), err)
		return ssoModels.Tokens{}, err
	}
	tokens := ssoModels.Tokens{
		AccessToken:  tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}
	return tokens, nil
}

func (c ClientSso) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	accessDetails, err := c.tk.ParseAccessDetails(token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("ParseAccessDetails token=%s", token), err)
		err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
		return false, err
	}

	_, err = c.rd.FetchAuth(accessDetails.TokenUuid)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("FetchAuth tokenUuid=%s", accessDetails.TokenUuid), err)
		return false, err
	}

	// todo: здесь через casbin делать проверку прав доступа
	//_, err := c.getTokenClaims(token)
	//if err != nil {
	//	err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
	//	return false, err
	//}
	return true, nil
}

func (c ClientSso) GetUserID(ctx context.Context, token string) (int, error) {
	//claims, err := c.getTokenClaims(token)
	//if err != nil {
	//	err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
	//	return 0, err
	//}
	//return claims.UserID, nil
	panic("implement me")
}

func checkUserPassword(user models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
