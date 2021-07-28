package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"
	"time"
)

type tokenService struct{}

func NewToken() *tokenService {
	return &tokenService{}
}

type TokenInterface interface {
	CreateToken(userID string) (*TokenDetails, error)
	ExtractTokenMetadata(tokenString string) (*AccessDetails, error)
}

var _ TokenInterface = &tokenService{}

func (t *tokenService) CreateToken(userID string) (*TokenDetails, error) {
	tokenDetails := &TokenDetails{}
	tokenDetails.AccessTokenExpires = time.Now().Add(time.Minute * 30).Unix()
	tokenDetails.AccessTokenUuid = uuid.New().String()

	tokenDetails.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshTokenUuid = tokenDetails.AccessTokenUuid + "++" + userID

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = tokenDetails.AccessTokenUuid
	atClaims["user_id"] = userID
	atClaims["exp"] = tokenDetails.AccessTokenExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	tokenDetails.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshTokenUuid = tokenDetails.AccessTokenUuid + "++" + userID

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetails.RefreshTokenUuid
	rtClaims["user_id"] = userID
	rtClaims["exp"] = tokenDetails.RefreshTokenExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	tokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return tokenDetails, nil
}

func TokenValid(tokenString string) error {
	token, err := verifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
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

func extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userID, userOK := claims["user_id"].(string)
		if !ok || !userOK {
			return nil, errors.New("unauthorized")
		} else {
			return &AccessDetails{
				TokenUuid: accessUuid,
				UserID:    userID,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func (t *tokenService) ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	acc, err := extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
