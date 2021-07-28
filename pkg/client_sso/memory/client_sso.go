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
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"os"
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

var signingKey = "merlinSigningKey"        //TODO cfg
var signingMethod = jwt.SigningMethodHS256 //TODO cfg

type tokenClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
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
	// todo: здесь же где-то должен делать маппинг ошибок
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
	saveErr := c.rd.CreateAuth(strconv.Itoa(user.ID), ts)
	if saveErr != nil {
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
	metadata, _ := c.tk.ExtractTokenMetadata(token)
	if metadata != nil {
		deleteErr := c.rd.DeleteTokens(metadata)
		if deleteErr != nil {
			deleteErr = wrapErr.NewWrapErr(fmt.Errorf("rd DeleteTokens"), deleteErr)
			return deleteErr
		}
	}
	return nil
}

func (c ClientSso) RefreshToken(ctx context.Context, refreshToken string) (ssoModels.Tokens, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("jwt parse"), err)
		return ssoModels.Tokens{}, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		err = wrapErr.NewWrapErr(fmt.Errorf("token is invalid or incorrect"), err)
		return ssoModels.Tokens{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = wrapErr.NewWrapErr(fmt.Errorf("token claims to jwt.MapClaims"), err)
		return ssoModels.Tokens{}, err
	}
	refreshUuid, ok := claims["refresh_uuid"].(string)
	if !ok {
		//
	}
	userId, roleOk := claims["user_id"].(string)
	if !roleOk {
		//
	}
	delErr := c.rd.DeleteRefresh(refreshUuid)
	if delErr != nil {
		//
	}
	ts, createErr := c.tk.CreateToken(userId)
	if createErr != nil {
		//
	}
	saveErr := c.rd.CreateAuth(userId, ts)
	if saveErr != nil {
		//
	}
	tokens := ssoModels.Tokens{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}
	return tokens, nil
}

func (c ClientSso) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	if err := auth.TokenValid(token); err != nil {
		err = wrapErr.NewWrapErr(errs.ErrIncorrectToken, err)
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
