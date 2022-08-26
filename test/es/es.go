package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"time"
)

var esAddr = "http://10.20.110.51:39200/"

type Employee struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

//创建索引
func create(client *elastic.Client) {
	//1.使用结构体方式存入到es里面
	e1 := Employee{"jane", "Smith", 20, "I like music", []string{"music"}}
	put, err := client.Index().Index("info").Type("employee").Id("1").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("indexed %v to index %s, type %s \n", put.Id, put.Index, put.Type)
}

func main() {
	fmt.Println("start main")
	var err error
	client, err := elastic.NewClient(
		elastic.SetURL(esAddr),
		elastic.SetBasicAuth("elasticbak", "Tsit@123"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	client.Start()
	info, code, err := client.Ping(esAddr).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Es return with code %v and version %s \n", code, info.Version.Number)
	esversionCode, err := client.ElasticsearchVersion(esAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("es version %s\n", esversionCode)
	//添加数据
	//create(client)
	IsRepExist(client)
	//restore(client)
}

func restore(client *elastic.Client) {
	restoreService := client.SnapshotRestore("backup", "snapshot_infobak")
	restoreBody := restoreService.BodyString("{\n    \"indices\": \"info\"\n}")
	restoreResponse, err := restoreBody.Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("restore 成功", restoreResponse.Accepted)
	//fmt.Println(restoreResponse.Snapshot.Snapshot)
}

//测试获取rep
func IsRepExist(client *elastic.Client) bool {
	response, err := client.SnapshotGetRepository("tesdddddddt").Do(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(len(response))
	if len(response) != 0 {
		return true
	}
	return false
}

func createrep(client *elastic.Client) {
	Snapshot := client.SnapshotCreateRepository("backup")
	Snapshot.BodyJson("{\n    \"type\": \"fs\", \n    \"settings\": {\n        \"location\": \"/usr/share/elasticsearch/bakfile\" \n    }\n}\n")
	res, err := Snapshot.Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Acknowledged)
}

func startbak(client *elastic.Client) {
	//创建快照
	SnapshotService := client.SnapshotCreate("backup", "snapshot_infobak")
	snapshotCreateResponse, err := SnapshotService.Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !*snapshotCreateResponse.Accepted {
		fmt.Println("备份失败")
	}
	fmt.Println(snapshotCreateResponse.Snapshot)
	fmt.Println("创建成功")
}

func getSnapShot(client *elastic.Client) {
	snapGetservice := client.SnapshotGet("backup")
	getResponse, err := snapGetservice.Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(getResponse.Snapshots)
	for _, snap := range getResponse.Snapshots {
		fmt.Println(snap)
	}
}
