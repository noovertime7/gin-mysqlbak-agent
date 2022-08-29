package elasticbak

import (
	"backupAgent/domain/config"
	"backupAgent/domain/pkg/log"
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"time"
)

type EsBaker interface {
	EsBakerInfo
	createRepository(ctx context.Context) error
	GetRepository(ctx context.Context) (elastic.SnapshotGetRepositoryResponse, error)
	CreateSnapshot(ctx context.Context, snap string) error
	GetSnapshot(ctx context.Context) (*elastic.SnapshotGetResponse, error)
	DeleteSnapshot(ctx context.Context, snap string) error
}

type EsBakerInfo interface {
	GetRepositoryDir() string
	GetRepositoryName() string
}

type EsHostInfo struct {
	Host     string
	UserName string
	Password string
}

func NewEsBaker(info *EsHostInfo) (EsBaker, error) {
	c, err := elastic.NewClient(
		elastic.SetURL(info.Host),
		elastic.SetBasicAuth(info.UserName, info.Password),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	e := &esBak{
		Client:         c,
		RepositoryName: config.GetStringConf("EsBackup", "RepositoryName"),
		RepositoryDir:  config.GetStringConf("EsBackup", "RepositoryDir"),
	}
	//初始化快照仓库
	if err := e.createRepository(context.TODO()); err != nil {
		return nil, errors.New("初始化仓库失败" + err.Error())
	}
	return e, nil
}

type esBak struct {
	Client         *elastic.Client
	RepositoryName string
	RepositoryDir  string
}

func (e *esBak) GetRepositoryName() string {
	return e.RepositoryName
}

func (e *esBak) GetRepositoryDir() string {
	return e.RepositoryDir
}

func (e *esBak) createRepository(ctx context.Context) error {
	//仓库存在不需要创建
	if e.IsRepositoryExist() {
		return nil
	}
	createRepository := e.Client.SnapshotCreateRepository(e.RepositoryName)
	params := fmt.Sprintf("{\n    \"type\": \"fs\", \n    \"settings\": {\n        \"location\": \"%s\" \n    }\n}", e.RepositoryDir)
	createRepository.BodyJson(params)
	data, err := createRepository.Do(ctx)
	if err != nil {
		return err
	}
	if data.Acknowledged {
		return nil
	}
	return errors.New("无报错但是Accepted返回不是true")
}

func (e *esBak) GetRepository(ctx context.Context) (elastic.SnapshotGetRepositoryResponse, error) {
	if !e.IsRepositoryExist() {
		return nil, errors.New("repository不存在")
	}
	return e.Client.SnapshotGetRepository(e.RepositoryName).Do(ctx)
}

func (e *esBak) CreateSnapshot(ctx context.Context, snap string) error {
	createResponse, err := e.Client.SnapshotCreate(e.RepositoryName, snap).Do(ctx)
	if err != nil {
		return err
	}
	if *createResponse.Accepted {
		return nil
	}
	return errors.New("无报错但是Accepted返回不是true")
}

func (e *esBak) GetSnapshot(ctx context.Context) (*elastic.SnapshotGetResponse, error) {
	return e.Client.SnapshotGet(e.RepositoryName).Do(ctx)
}

func (e *esBak) DeleteSnapshot(ctx context.Context, snap string) error {
	deleteResponse, err := e.Client.SnapshotDelete(e.RepositoryName, snap).Do(ctx)
	if err != nil {
		return err
	}
	if deleteResponse.Acknowledged {
		return nil
	}
	return errors.New("无报错但是Accepted返回不是true")
}

func (e *esBak) IsRepositoryExist() bool {
	response, err := e.Client.SnapshotGetRepository(e.RepositoryName).Do(context.Background())
	if err != nil {
		log.Logger.Warn(err)
		return false
	}
	if len(response) != 0 {
		return true
	}
	return false
}
