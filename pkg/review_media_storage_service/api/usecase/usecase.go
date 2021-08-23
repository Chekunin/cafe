package usecase

import (
	"cafe/pkg/media_storage"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
)

type Usecase struct {
	mediaStorage media_storage.IMediaStorage
}

func NewUsecase(mediaStorage media_storage.IMediaStorage) *Usecase {
	return &Usecase{
		mediaStorage: mediaStorage,
	}
}

func (u Usecase) GetStreamMedia(ctx context.Context, path string) (io.ReadCloser, string, error) {
	readCloser, contentType, err := u.mediaStorage.GetStream(ctx, path)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("mediaStorage GetStream path=%s", path), err)
		return nil, "", err
	}
	return readCloser, contentType, nil
}

func (u Usecase) PutStreamMedia(ctx context.Context, path string, reader io.Reader) (media_storage.Object, error) {
	object, err := u.mediaStorage.Put(ctx, path, reader)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("mediaStorage Put path=%s", path), err)
		return media_storage.Object{}, err
	}
	return object, nil
}
