package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type service struct {
	client *redis.Client
}

var _ AuthInterface = &service{}

func NewAuth(client *redis.Client) *service {
	return &service{client: client}
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessTokenUuid     string
	RefreshTokenUuid    string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

type AccessDetails struct {
	TokenUuid string
	UserID    string
}

//Save token metadata to Redis
func (s service) CreateAuth(userId string, td *TokenDetails) error {
	at := time.Unix(td.AccessTokenExpires, 0)
	rt := time.Unix(td.RefreshTokenExpires, 0)
	now := time.Now()

	atCreated, err := s.client.Set(context.TODO(), td.AccessTokenUuid, userId, at.Sub(now)).Result()
	if err != nil {
		return nil
	}
	rtCreated, err := s.client.Set(context.TODO(), td.RefreshTokenUuid, userId, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

//Check the metadata saved
func (s service) FetchAuth(tokenUuid string) (string, error) {
	userID, err := s.client.Get(context.TODO(), tokenUuid).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (s service) DeleteRefresh(refreshUuid string) error {
	deleted, err := s.client.Del(context.TODO(), refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}

//Once a user row in the token table
func (s service) DeleteTokens(authD *AccessDetails) error {
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UserID)
	deletedAt, err := s.client.Del(context.TODO(), authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	deletedRt, err := s.client.Del(context.TODO(), refreshUuid).Result()
	if err != nil {
		return err
	}
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
