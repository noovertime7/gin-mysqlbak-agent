package alioss

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewClient(filename, Endpoint, AccessKey, SecretKey, BucketName, Directory string) (*AliOss, error) {
	client, err := oss.New(Endpoint, AccessKey, SecretKey)
	if err != nil {
		log.Logger.Error("Ali OSS Error:", err)
		return nil, err
	}
	return &AliOss{
		client:     client,
		fileName:   filename,
		BucketName: BucketName,
		Dir:        Directory,
	}, nil
}

type AliOss struct {
	client     *oss.Client
	fileName   string
	BucketName string
	Dir        string
}

func (a *AliOss) AliOssUploadFile() error {
	if !a.isExists() {
		if err := a.createBucket(); err != nil {
			return err
		}
	}
	bucket, err := a.client.Bucket(a.BucketName)
	if err != nil {
		return err
	}
	file := strings.Split(a.fileName, "/")[len(strings.Split(a.fileName, "/"))-1] //需要处理一下拿到文件名
	return bucket.PutObjectFromFile(a.Dir+file, a.fileName)
}

func (a *AliOss) Remove() error {
	bucket, err := a.client.Bucket(a.BucketName)
	if err != nil {
		return err
	}
	filename, _ := filepath.Split(a.fileName)
	path := "/" + a.Dir + "/" + filename
	return bucket.DeleteObject(path)
}

func (a *AliOss) isExists() bool {
	ok, _ := a.client.IsBucketExist(a.BucketName)
	return ok
}

func (a *AliOss) createBucket() error {
	return a.client.CreateBucket(a.BucketName)
}
