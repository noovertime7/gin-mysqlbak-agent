package minio

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"path/filepath"
)

type Client struct {
	client         *minio.Client
	bucketName     string
	targetFilePath string
	Dir            string
}

func NewClient(endpoint, accessKeyID, secretAccessKey, bucketName, dir, filepath string) (*Client, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		client:         minioClient,
		bucketName:     bucketName, //目标bucket
		targetFilePath: filepath,
		Dir:            dir,
	}, nil
}

func (c *Client) isExists() bool {
	ok, _ := c.client.BucketExists(context.Background(), c.bucketName)
	return ok
}

// checkBucket 检测目标bucket是否存在，不存在就创建一个
func (c *Client) checkBucket() {
	if !c.isExists() {
		err2 := c.client.MakeBucket(context.Background(), c.bucketName, minio.MakeBucketOptions{Region: "cn-north-1", ObjectLocking: false})
		if err2 != nil {
			logger.Error("MakeBucket error ", err2)
			return
		}
		logger.Info("minio创建bucket %s\n", c.bucketName)
	}
}

func (c *Client) UploadFile() error {
	c.checkBucket()
	_, filename := filepath.Split(c.targetFilePath)
	dirname := "/" + c.Dir + "/" + filename
	_, err := c.client.FPutObject(context.Background(), c.bucketName, dirname, c.targetFilePath, minio.PutObjectOptions{})
	if err != nil {
		logger.Error("上传失败 ", err)
		return err
	}
	return nil
}

func (c *Client) Remove() error {
	_, filename := filepath.Split(c.targetFilePath)
	path := "/" + c.Dir + "/" + filename
	logger.Infof("minio handler delete file %s", path)
	return c.client.RemoveObject(context.Background(), c.bucketName, path, minio.RemoveObjectOptions{GovernanceBypass: true})
}
