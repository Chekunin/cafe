package media_storage

import (
	"context"
	"io"
	"os"
	"time"
)

type IMediaStorage interface {
	Get(ctx context.Context, path string) (*os.File, error)
	GetStream(ctx context.Context, path string) (io.ReadCloser, string, error)
	Put(ctx context.Context, path string, reader io.Reader) (Object, error)
	Delete(ctx context.Context, path string) error
	List(ctx context.Context, path string) ([]Object, error)
}

// Object content object
type Object struct {
	Path         string
	Name         string
	LastModified *time.Time
}
