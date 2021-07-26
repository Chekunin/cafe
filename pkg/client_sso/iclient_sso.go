package client_sso

import "context"

type IClientSso interface {
	Login(ctx context.Context, userName string, password string) (string, error)
	CheckPermission(ctx context.Context, method, path, token string) (bool, error)
	GetUserID(ctx context.Context, token string) (int, error)
}
