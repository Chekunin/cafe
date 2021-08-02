package common

import "context"

type contextKey int

const (
	ContextKeyRequestId contextKey = iota
	ContextKeyToken     contextKey = iota
)

func FromContextRequestId(ctx context.Context) (string, bool) {
	u, ok := ctx.Value(ContextKeyRequestId).(string)
	return u, ok
}

func FromContextToken(ctx context.Context) (string, bool) {
	u, ok := ctx.Value(ContextKeyRequestId).(string)
	return u, ok
}
