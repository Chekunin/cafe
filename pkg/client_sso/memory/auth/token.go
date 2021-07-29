package auth

import (
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

var signingMethod = jwt.SigningMethodHS256
var accessTokenSecretKey = "merlinSigningKey"
var refreshTokenSecretKey = "merlinSigningKey"

//var accessTokenExpires = time.Minute * 30
//var refreshTokenExpires = time.Hour * 24 * 7
var accessTokenExpires = time.Minute * 1
var refreshTokenExpires = time.Minute * 5

type tokenService struct{}

func NewToken() *tokenService {
	return &tokenService{}
}

type TokenInterface interface {
	CreateToken(userID string) (TokenDetails, error)
	ParseAccessDetails(tokenString string) (AccessDetails, error)
}

var _ TokenInterface = &tokenService{}

type TokenClaims struct {
	UserID    string `json:"user_id"`
	TokenUuid string `json:"token_uuid"`
	jwt.StandardClaims
}

func (t *tokenService) CreateToken(userID string) (TokenDetails, error) {
	tokenDetails := TokenDetails{}

	tokenDetails.AccessTokenExpires = time.Now().Add(accessTokenExpires).Unix()
	tokenDetails.AccessTokenUuid = uuid.New().String()
	atClaims := &TokenClaims{
		UserID:    userID,
		TokenUuid: tokenDetails.AccessTokenUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenDetails.AccessTokenExpires,
		},
	}
	at := jwt.NewWithClaims(signingMethod, atClaims)
	var err error
	tokenDetails.AccessToken, err = at.SignedString([]byte(accessTokenSecretKey))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("at SignedString"), err)
		return TokenDetails{}, err
	}

	tokenDetails.RefreshTokenExpires = time.Now().Add(refreshTokenExpires).Unix()
	tokenDetails.RefreshTokenUuid = tokenDetails.AccessTokenUuid + "++" + userID
	rtClaims := &TokenClaims{
		UserID:    userID,
		TokenUuid: tokenDetails.RefreshTokenUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenDetails.RefreshTokenExpires,
		},
	}
	rt := jwt.NewWithClaims(signingMethod, rtClaims)
	tokenDetails.RefreshToken, err = rt.SignedString([]byte(refreshTokenSecretKey))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("rt SignedString"), err)
		return TokenDetails{}, err
	}
	return tokenDetails, nil
}

func (t *tokenService) ParseAccessDetails(tokenString string) (AccessDetails, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return AccessDetails{}, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		err := wrapErr.NewWrapErr(fmt.Errorf("token is invalid"), nil)
		return AccessDetails{}, err
	}
	return AccessDetails{
		TokenUuid: claims.TokenUuid,
		UserID:    claims.UserID,
	}, nil
}

// принимает строку с token-ом, возвращает его распарсенный объект
// отсутствие ошибки говорит, что с токеном всё норм
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessTokenSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
