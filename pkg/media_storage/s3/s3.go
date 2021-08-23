package s3

import (
	"bytes"
	"cafe/pkg/media_storage"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Client struct {
	*s3.Client
	Config *Config
}

type Config struct {
	AccessID         string `yaml:"access_id"`
	AccessKey        string `yaml:"access_key"`
	Region           string `yaml:"region"`
	Bucket           string `yaml:"bucket"`
	S3ForcePathStyle bool   `yaml:"s3_force_path_style"`
}

func New(cfg *Config) *Client {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessID, cfg.AccessKey, "")),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("aws config LoadDefaultConfig"), err)
		panic(err)
	}

	client := s3.NewFromConfig(awsConfig)

	return &Client{
		Client: client,
		Config: cfg,
	}
}

func (c Client) Get(ctx context.Context, path string) (*os.File, error) {
	readCloser, _, err := c.GetStream(ctx, path)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("GetStream path=%s", path), err)
		return nil, err
	}

	ext := filepath.Ext(path)
	pattern := fmt.Sprintf("s3*%s", ext)

	var file *os.File
	file, err = os.CreateTemp("tmp", pattern)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("os.CreateTemp pattern=%s", pattern), err)
		return nil, err
	}
	fmt.Printf("file name = %s\n", file.Name())

	defer readCloser.Close()
	_, err = io.Copy(file, readCloser)
	file.Seek(0, 0)

	return file, err

}

// GetStream get file as stream
func (c Client) GetStream(ctx context.Context, path string) (io.ReadCloser, string, error) {
	getResponse, err := c.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.Config.Bucket),
		//Key:    aws.String(c.ToRelativePath(path)),
		Key: aws.String(path),
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("GetObject bucket=%s key=%s", c.Config.Bucket, path), err)
		return nil, "", err
	}

	return getResponse.Body, *getResponse.ContentType, nil
}

func (c Client) Put(ctx context.Context, urlPath string, reader io.Reader) (media_storage.Object, error) {
	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("ioutil ReadAll"), err)
		return media_storage.Object{}, err
	}
	fileType := mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(c.Config.Bucket), // required
		Key:    aws.String(urlPath),         // required
		//ACL:    aws.String(client.Config.ACL),
		Body: bytes.NewReader(buffer),
		//ContentLength: aws.Int64(int64(len(buffer))),
		ContentType: aws.String(fileType),
	}
	//if c.Config.CacheControl != "" {
	//	params.CacheControl = aws.String(client.Config.CacheControl)
	//}

	_, err = c.Client.PutObject(ctx, params)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("c client PutObject params=%+v", params), err)
		return media_storage.Object{}, err
	}

	now := time.Now()
	return media_storage.Object{
		Path:         urlPath,
		Name:         filepath.Base(urlPath),
		LastModified: &now,
	}, err
}

func (c Client) Delete(ctx context.Context, path string) error {
	panic("implement me")
}

func (c Client) List(ctx context.Context, path string) ([]media_storage.Object, error) {
	var objects []media_storage.Object
	var prefix string

	if path != "" {
		prefix = strings.Trim(path, "/") + "/"
	}

	listObjectsResponse, err := c.Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.Config.Bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("client ListObjectsV2"), err)
		return nil, err
	}

	for _, content := range listObjectsResponse.Contents {
		objects = append(objects, media_storage.Object{
			Path:         *content.Key,
			Name:         filepath.Base(*content.Key),
			LastModified: content.LastModified,
		})
	}

	return objects, nil
}
