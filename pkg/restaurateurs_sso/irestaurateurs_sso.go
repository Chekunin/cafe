package restaurateurs_sso

import (
	"cafe/pkg/restaurateurs_sso/models"
	"context"
)

type IRestaurateursSso interface {
	Login(ctx context.Context, email string, password string) (models.Tokens, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error)
	CheckPermission(ctx context.Context, method, path, token string) (models.RespCheckPermission, error)
	GetRestaurateurID(ctx context.Context, token string) (int, error)
}
