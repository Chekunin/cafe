package client_sso

import (
	"cafe/pkg/client_sso/models"
	"context"
)

type IClientSso interface {
	Login(ctx context.Context, userName string, password string) (models.Tokens, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error)
	CheckPermission(ctx context.Context, method, path, token string) (models.RespCheckPermission, error)
	GetUserID(ctx context.Context, token string) (int, error)
}
