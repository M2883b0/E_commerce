package data

import (
	"content_manage/internal/conf"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
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
	database_conf := os.Getenv("MYSQL_ADDR")
	if database_conf == "" {
		database_conf = c.GetDatabase().GetSource()
	}
	mysqlDB, er := gorm.Open(mysql.Open(database_conf))
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
	var EsCfg elasticsearch.Config
	if os.Getenv("ELASTIC_USERNAME") == "" {
		////ES初始化
		EsCfg = elasticsearch.Config{
			Addresses: []string{"https://127.0.0.1:9200"}, // 从配置获取ES地址
			Username:  "elastic",
			Password:  "6yrXKjLlSXY4V_lCnMss",
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 禁用TLS验证
				},
			},
		}
	} else {
		EsCfg = elasticsearch.Config{
			Addresses: []string{os.Getenv("ELASTIC_ADDR")}, // 从配置获取ES地址
			Username:  os.Getenv("ELASTIC_USERNAME"),
			Password:  os.Getenv("ELASTIC_PASSWORD"),
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // 禁用TLS验证
				},
			},
		}
	}

	// 创建ES客户端
	esClient, err := elasticsearch.NewClient(EsCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create ES client: %v", err)
	}
	data := &Data{
		db: mysqlDB,
		es: esClient,
	}

	// 初始化 Elasticsearch  （同步数据也行）
	if err := data.SyncToElasticsearch(context.Background()); err != nil {
		fmt.Printf("Es初始化失败: Failed to init index data to Elasticsearch: %v", err)
	}

	return data, cleanup, nil
}

// 初始化Elasticsearch的Index
func (d *Data) SyncToElasticsearch(ctx context.Context) error {
	// 检查 Elasticsearch 是否已经初始化了名为“products”的Index
	// 1. 检查索引是否存在
	existsReq := esapi.IndicesExistsRequest{
		Index: []string{"products"},
	}

	res, err := existsReq.Do(ctx, d.es)
	if err != nil {
		return fmt.Errorf("检查索引存在失败: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Printf("====== Elasticsearch Index already has data, skipping init ======\n")
		return nil // 索引已存在
	}
	mapping := `
{
 "mappings": {
   "properties": {
     "id": {
       "type": "long"
     },
     "title": {
       "type": "text"
     },
     "description": {
       "type": "text"
     },
     "categories": {
       "type": "text"
     }
   }
 }
}`
	// 3. 创建初始化的index索引
	createReq := esapi.IndicesCreateRequest{
		Index: "products",
		Body:  strings.NewReader(mapping),
	}

	createRes, err := createReq.Do(ctx, d.es)
	if err != nil {
		return fmt.Errorf("创建索引失败: %v", err)
	}
	defer createRes.Body.Close()

	if createRes.IsError() {
		return fmt.Errorf("索引创建错误: %s", createRes.String())
	}
	fmt.Printf("Index创建成功，Es初始化完成\n")

	return nil
	//如果不存在名为“products”的Index，则创建它

	//把mysql的全部数据，导入Es中
	//// 检查 Elasticsearch 中是否已经有数据
	//res, err := d.es.Count(
	//	d.es.Count.WithIndex("products"),
	//)
	//if err != nil {
	//	return fmt.Errorf("failed to count documents in Elasticsearch: %v", err)
	//}
	//defer res.Body.Close()
	//
	//var countResponse map[string]interface{}
	//if err := json.NewDecoder(res.Body).Decode(&countResponse); err != nil {
	//	return fmt.Errorf("failed to decode Elasticsearch count response: %v", err)
	//}
	//
	//count := int(countResponse["count"].(float64))
	//if count > 0 {
	//	fmt.Println("====== Elasticsearch already has data, skipping sync ======")
	//	return nil
	//}
	//
	//// 否则需要构建，从 MySQL 中获取所有商品数据
	//var details []*ContentDetail
	//if err := d.db.Find(&details).Error; err != nil {
	//	return fmt.Errorf("failed to fetch data from MySQL: %v", err)
	//}
	//
	//// 将数据同步到 Elasticsearch
	//for _, detail := range details {
	//	doc := map[string]interface{}{
	//		"id":          detail.ID,
	//		"title":       detail.Title,
	//		"description": detail.Description,
	//		"categories":  detail.Categories,
	//	}
	//
	//	docJSON, err := json.Marshal(doc)
	//	if err != nil {
	//		fmt.Printf("Failed to marshal document: %v", err)
	//		continue
	//	}
	//
	//	req := esapi.IndexRequest{
	//		Index:      "products",
	//		DocumentID: fmt.Sprintf("%d", detail.ID),
	//		Body:       strings.NewReader(string(docJSON)),
	//		Refresh:    "true",
	//	}
	//
	//	res, err := req.Do(ctx, d.es)
	//	if err != nil {
	//		fmt.Printf("Failed to index document: %v", err)
	//		continue
	//	}
	//	defer res.Body.Close()
	//
	//	if res.IsError() {
	//		fmt.Printf("Error indexing document ID=%d: %s", detail.ID, res.String())
	//	} else {
	//		fmt.Printf("Indexed document ID=%d", detail.ID)
	//	}
	//}
	//return nil
}
