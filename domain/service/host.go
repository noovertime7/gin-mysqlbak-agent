package service

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/proto/backupAgent/host"
	"backupAgent/proto/backupAgent/task"
	"context"
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/olivere/elastic"
	"time"
)

type HostService struct{}

func (h *HostService) HostAdd(ctx context.Context, hostInfo *host.HostAddInput) error {
	//进行主机检测避免添加无用信息
	if err := HostPingCheck(hostInfo.UserName, hostInfo.Password, hostInfo.Host, "", pkg.HostType(hostInfo.Type)); err != nil {
		log.Logger.Error("agent添加主机检测失败", err)
		return err
	}
	hostDB := &dao.HostDatabase{
		Host:       hostInfo.Host,
		User:       hostInfo.UserName,
		Password:   hostInfo.Password,
		Content:    hostInfo.Content,
		HostStatus: 1,
		IsDeleted:  0,
		Type:       hostInfo.Type,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return hostDB.Save(ctx, database.Gorm)
}

func (h *HostService) HostDelete(ctx context.Context, hid int64) error {
	hostDB := &dao.HostDatabase{Id: hid}
	hostInfo, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return err
	}
	if hostInfo.Id == 0 {
		return errors.New("主机不存在,请检查id是否正确")
	}
	hostInfo.IsDeleted = 1
	return hostInfo.Save(ctx, database.Gorm)
}

func (h *HostService) HostUpdate(ctx context.Context, hostInfo *host.HostUpdateInput) error {
	//进行主机检测避免添加无用信息
	if err := HostPingCheck(hostInfo.UserName, hostInfo.Password, hostInfo.Host, "", pkg.HostType(hostInfo.Type)); err != nil {
		log.Logger.Error("agent添加主机检测失败", err)
		return err
	}
	hostDB := &dao.HostDatabase{
		Id:         hostInfo.ID,
		Host:       hostInfo.Host,
		User:       hostInfo.UserName,
		Password:   hostInfo.Password,
		Content:    hostInfo.Content,
		HostStatus: 1,
		Type:       hostInfo.Type,
	}
	return hostDB.Save(ctx, database.Gorm)
}

func (h *HostService) GetHostList(ctx context.Context, hostInfo *host.HostListInput) (*host.HostListOutPut, error) {
	hostDB := &dao.HostDatabase{}
	hostList, total, err := hostDB.PageList(ctx, database.Gorm, hostInfo)
	if err != nil {
		return nil, err
	}
	var outList []*host.ListItem
	taskDB := &dao.TaskInfo{}
	for _, listIterm := range hostList {
		_, total, err := taskDB.PageList(ctx, database.Gorm, &task.TaskListInput{HostID: listIterm.Id})
		if err != nil {
			return nil, err
		}
		outIterm := &host.ListItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			UserName:   listIterm.User,
			Password:   listIterm.Password,
			HostStatus: listIterm.HostStatus,
			Content:    listIterm.Content,
			TaskNum:    total,
			Type:       listIterm.Type,
			CreateAt:   listIterm.CreatedAt.Format("2006年01月02日15:04:01"),
			UpdateAt:   listIterm.UpdatedAt.Format("2006年01月02日15:04:01"),
		}
		outList = append(outList, outIterm)
	}
	return &host.HostListOutPut{
		Total:    total,
		ListItem: outList,
	}, nil
}

func HostPingCheck(User, Password, Host, DBName string, hostType pkg.HostType) error {
	switch hostType {
	case pkg.MysqlHost:
		if err := MysqlHostCheck(User, Password, Host, DBName); err != nil {
			return err
		}
	case pkg.ElasticHost:
		if err := EsHostCheck(Host, User, Password); err != nil {
			return err
		}
	}
	return nil
}

func MysqlHostCheck(User, Password, Host, DBName string) error {
	if DBName == "" {
		DBName = "mysql"
	}
	en, err := xorm.NewEngine("mysql", User+":"+Password+"@tcp("+Host+")/"+DBName+"?charset=utf8&parseTime=true")
	defer en.Close()
	if err != nil {
		log.Logger.Errorf("创建数据库连接失败:%s", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = en.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func EsHostCheck(host, user, password string) error {
	if _, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetBasicAuth(user, password),
		elastic.SetSniff(false)); err != nil {
		return err
	}
	return nil
}
