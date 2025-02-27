package data

import (
	"content_manage/internal/conf"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContentRepo)

// Data .
type Data struct {
	db *gorm.DB
	es *elasticsearch.Client //新增ElasticSearch
	//gorse *client.GorseClient   // 新增Gorse
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	//mysql初始化
	mysqlDB, er := gorm.Open(mysql.Open(c.GetDatabase().GetSource()))
	if er != nil {
		panic(er)
	}
	mysqlDB.AutoMigrate(&ContentDetail{}) //自动迁移，自动创建表，默认蛇行负数
	//拿到mysqlDB的实例
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2

	//ES初始化
	var esConfig conf.Elasticsearch
	esCfg := elasticsearch.Config{
		Addresses: esConfig.Addresses, // 从配置获取ES地址
		Username:  esConfig.Username,
		Password:  esConfig.Password,
	}
	// 创建ES客户端
	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create ES client: %v", err)
	}
	data := &Data{
		db: mysqlDB,
		es: esClient,
	}
	//// 同步数据到 Elasticsearch
	//if err := data.SyncToElasticsearch(context.Background()); err != nil {
	//	fmt.Printf("Failed to sync data to Elasticsearch: %v", err)
	//}

	return data, cleanup, nil
}

// 将 MySQL 中的商品数据同步到 Elasticsearch 索引中
func (d *Data) SyncToElasticsearch(ctx context.Context) error {
	// 检查 Elasticsearch 是否已经初始化了名为“products”的Index
	//如果不存在名为“products”的Index，则创建它

	// 检查 Elasticsearch 中是否已经有数据
	res, err := d.es.Count(
		d.es.Count.WithIndex("products"),
	)
	if err != nil {
		return fmt.Errorf("failed to count documents in Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	var countResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&countResponse); err != nil {
		return fmt.Errorf("failed to decode Elasticsearch count response: %v", err)
	}

	count := int(countResponse["count"].(float64))
	if count > 0 {
		fmt.Println("====== Elasticsearch already has data, skipping sync ======")
		return nil
	}

	// 否则需要构建，从 MySQL 中获取所有商品数据
	var details []*ContentDetail
	if err := d.db.Find(&details).Error; err != nil {
		return fmt.Errorf("failed to fetch data from MySQL: %v", err)
	}

	// 将数据同步到 Elasticsearch
	for _, detail := range details {
		doc := map[string]interface{}{
			"id":          detail.ID,
			"title":       detail.Title,
			"description": detail.Description,
			"categories":  detail.Categories,
		}

		docJSON, err := json.Marshal(doc)
		if err != nil {
			fmt.Printf("Failed to marshal document: %v", err)
			continue
		}

		req := esapi.IndexRequest{
			Index:      "products",
			DocumentID: fmt.Sprintf("%d", detail.ID),
			Body:       strings.NewReader(string(docJSON)),
			Refresh:    "true",
		}

		res, err := req.Do(ctx, d.es)
		if err != nil {
			fmt.Printf("Failed to index document: %v", err)
			continue
		}
		defer res.Body.Close()

		if res.IsError() {
			fmt.Printf("Error indexing document ID=%d: %s", detail.ID, res.String())
		} else {
			fmt.Printf("Indexed document ID=%d", detail.ID)
		}
	}

	return nil
}
