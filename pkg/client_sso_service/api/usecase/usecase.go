package usecase

import (
	"cafe/pkg/client_sso"
	"cafe/pkg/client_sso/models"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

type Usecase struct {
	clientSso client_sso.IClientSso
}

func NewUsecase(clientSso client_sso.IClientSso) *Usecase {
	return &Usecase{clientSso: clientSso}
}

func (u Usecase) Login(ctx context.Context, userName string, password string) (models.Tokens, error) {
	tokens, err := u.clientSso.Login(ctx, userName, password)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso login"), err)
		return models.Tokens{}, err
	}
	return tokens, nil
}

func (u Usecase) Logout(ctx context.Context, token string) error {
	if err := u.clientSso.Logout(ctx, token); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso logout"), err)
		return err
	}
	return nil
}

func (u Usecase) RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error) {
	tokens, err := u.clientSso.RefreshToken(ctx, refreshToken)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso RefreshToken"), err)
		return models.Tokens{}, err
	}
	return tokens, nil
}

func (u Usecase) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	ok, err := u.clientSso.CheckPermission(ctx, method, path, token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso CheckPermission"), err)
		return false, err
	}
	return ok, nil
}

func (u Usecase) GetUserID(ctx context.Context, token string) (int, error) {
	userID, err := u.clientSso.GetUserID(ctx, token)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("clientSso GetUserID"), err)
		return 0, err
	}
	return userID, nil
}
