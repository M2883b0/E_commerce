package data

import (
	"content_manage/internal/biz"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	gorm.Model
	Title       string  `json:"column:title"`
	Description string  `json:"column:description"`
	Picture_url string  `json:"column:picture_url"`
	Price       float32 `json:"column:price"`
	Quantity    uint32  `json:"column:quantity"`
	Categories  string  `json:"column:categories"`
}

func (c ContentDetail) TableName() string {
	return "goods.product" //数据库的表名
}

type EsDetail struct {
	Id          int64  `json:"column:id"`
	Title       string `json:"column:title"`
	Description string `json:"column:description"`
	Categories  string `json:"column:categories"`
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	//// 开启事务
	//tx := c.data.db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//c.log.Infof("contentRepo Create context = %+v", content)
	////把字符串数组，转成字符串
	//categoriesStr := strings.Join(content.Categories, ",")
	//detail := ContentDetail{
	//	Title:       content.Title,
	//	Description: content.Description,
	//	Picture_url: content.Picture_url,
	//	Price:       content.Price,
	//	Quantity:    content.Quantity,
	//	Categories:  categoriesStr,
	//}
	//if err := tx.Create(&detail).Error; err != nil {
	//	tx.Rollback()
	//	c.log.Errorf("Mysql content create error = %v", err)
	//	return err
	//}
	////双写，写入数据库的同时，同步到ElasticSearch
	//var esdetail []*EsDetail
	//esdetail = append(esdetail, &EsDetail{
	//	Id:          int64(detail.ID),
	//	Title:       detail.Title,
	//	Description: detail.Description,
	//	Categories:  detail.Categories,
	//})
	//err := c.BatchUpSertToEs(ctx, esdetail) //调用UpSet方法（没有就创建，有就更新）
	//if err != nil {
	//	tx.Rollback() //Es失败时，回滚Mysql数据库
	//	c.log.Errorf("Es content create %v error = %v", esdetail, err)
	//	return err
	//}
	//// 提交事务
	//return tx.Commit().Error

	//只mysql操作，使用Canal检测mysql数据的变动
	c.log.Infof("新增商品请求 = %+v", content)
	categoriesStr := strings.Join(content.Categories, ",")
	detail := ContentDetail{
		Title:       content.Title,
		Description: content.Description,
		Picture_url: content.Picture_url,
		Price:       content.Price,
		Quantity:    content.Quantity,
		Categories:  categoriesStr,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Infof("商品新增错误 = %+v", err)
		return err
	}
	//立马更新商品的图片url名字，为ID.jpg
	db.Model(&detail).Where("id = ?", detail.ID).Update("picture_url", fmt.Sprintf("%d.jpg", detail.ID))
	return nil
}

// 暂时不考虑，商品内容更新后，Es的变更
func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	c.log.Infof("商品更新请求 = %+v", id)
	db := c.data.db
	//把字符串数组，转成字符串
	categoriesStr := strings.Join(content.Categories, ",")
	detail := ContentDetail{
		Title:       content.Title,
		Description: content.Description,
		Picture_url: content.Picture_url,
		Price:       content.Price,
		Quantity:    content.Quantity,
		Categories:  categoriesStr,
	}
	if err := db.Where("id = ?", id).
		Updates(&detail).Error; err != nil {
		c.log.Infof("商品更新错误 = %+v", err)
		return err
	}
	return nil
}

// Es批量更新或插入
func (c *contentRepo) BatchUpSertToEs(ctx context.Context, data []*EsDetail) error {
	if len(data) == 0 {
		return nil
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: c.data.es,
		Index:  "products",
	})
	if err != nil {
		return err
	}

	for _, d := range data {
		v, err := json.Marshal(d)
		if err != nil {
			return err
		}

		payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "update",
			DocumentID: fmt.Sprintf("%d", d.Id), //新增的商品id，作为id
			Body:       strings.NewReader(payload),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			},
		})
		if err != nil {
			return err
		}
	}

	return bi.Close(ctx)
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.Infof("商品不存在错误 = %+v", err)
		return false, err
	}
	return true, nil
}

// 不考虑删除商品信息，Es的同步问题
func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	c.log.Infof("商品删除请求 = %+v", id)
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).Delete(&ContentDetail{}).Error
	if err != nil {
		c.log.Infof("商品删除错误 = %+v", err)
		return err
	}
	return nil
}

func (c *contentRepo) Get(ctx context.Context, ids []int64) ([]*biz.Content, error) {
	c.log.Infof("商品查询请求 = %+v", ids)
	db := c.data.db
	var details []*ContentDetail
	//根据ids列表，一次查询全部的商品信息
	if err := db.Where("id in ?", ids).Find(&details).Error; err != nil {
		c.log.Infof("商品查询错误 = %+v", err)
		return nil, err
	}
	var contents []*biz.Content
	for _, detail := range details {
		contents = append(contents, &biz.Content{
			ID:          int64(detail.ID),
			Title:       detail.Title,
			Description: detail.Description,
			Picture_url: detail.Picture_url,
			Price:       detail.Price,
			Quantity:    detail.Quantity,
			Categories:  strings.Split(detail.Categories, ","),
		})
	}
	return contents, nil
}

// 商品库存变动，is_add为True，执行增加，is_add为False，执行减的操作。注意减少，会不会减少为0
func (c *contentRepo) UpdateQuantity(ctx context.Context, quantity_list []*biz.QuantityDetail) (bool, error) {
	db := c.data.db
	// 开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, q := range quantity_list {
		id := q.ID
		is_add := q.Is_add
		quantity := q.Quantity
		var detail ContentDetail
		if err := tx.Where("id = ?", id).First(&detail).Error; err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return false, fmt.Errorf("商品不存在: %v", err)
			}
			c.log.WithContext(ctx).Errorf("查询商品失败: %v", err)
			return false, err
		}

		if is_add == true {
			// 增加库存
			detail.Quantity += uint32(quantity)
		} else if is_add == false {
			// 减少库存，确保不会小于0
			if detail.Quantity < uint32(quantity) {
				tx.Rollback()
				return false, fmt.Errorf("库存不足，无法减少: 当前库存=%d, 请求减少=%d", detail.Quantity, quantity)
			}
			detail.Quantity -= uint32(quantity)
		} else {
			tx.Rollback()
			return false, fmt.Errorf("无效的is_add值，1-为增加库存，0-为减少库存: %d", is_add)
		}

		if err := tx.Save(&detail).Error; err != nil {
			tx.Rollback()
			c.log.WithContext(ctx).Errorf("更新库存失败: %v", err)
			return false, err
		}

	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.log.WithContext(ctx).Errorf("提交事务失败: %v", err)
		return false, err
	}

	return true, nil
}

// 搜索框，搜索商品，接入ElasticSearch + Kibana（查找）
func (c *contentRepo) Find(ctx context.Context, search string, in_page, in_pageSize int32) ([]*biz.Content, int64, error) {
	c.log.Infof("商品搜索请求 = %+v", search)

	// 构造模糊查询条件
	//query := c.data.db.Model(&ContentDetail{})
	//query = query.Where("title = %?%", search)
	//query = query.Where("description = %?%", search)
	//query = query.Where("categories = %?%", search)
	//
	//// 总数
	//var total int64
	//if err := query.Count(&total).Error; err != nil {
	//	return nil, 0, err
	//}
	////设置默认页大小
	//var page, pageSize = 1, 10
	//if in_page > 0 {
	//	page = int(in_page)
	//}
	//if in_pageSize > 0 {
	//	pageSize = int(in_pageSize)
	//}
	//offset := (page - 1) * pageSize
	////进行数据库查找
	//var results []*ContentDetail
	//if err := query.Offset(offset).Limit(pageSize).
	//	Find(&results).Error; err != nil {
	//	c.log.WithContext(ctx).Errorf("content find error = %v", err)
	//	return nil, 0, err
	//}
	//var contents []*biz.Content
	////将数据库查找的结构，映射到biz.Content定义的结构
	//for _, r := range results {
	//	contents = append(contents, &biz.Content{
	//		ID:          int64(r.ID),
	//		Title:       r.Title,
	//		Description: r.Description,
	//		Picture_url: r.Picture_url,
	//		Price:       r.Price,
	//		Quantity:    r.Quantity,
	//		Categories:  strings.Split(r.Categories, ","), //字符串转数组
	//	})
	//}
	//return contents, total, nil

	//=============================ES=============================
	//// 设置默认页大小
	var page, pageSize = 1, 10
	if in_page > 0 {
		page = int(in_page)
	}
	if in_pageSize > 0 {
		pageSize = int(in_pageSize)
	}
	offset := (page - 1) * pageSize

	fmt.Printf("用户的搜索内容是%+v", search)

	// 构造 Elasticsearch 查询条件
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  search,
				"fields": []string{"column:title", "column:description", "column:categories"},
				//"fields": []string{"title^3", "description^2", "categories"},
			},
		},
		"from": offset,
		"size": pageSize,
		"sort": []map[string]interface{}{
			{"_score": "desc"}, // 按相关度排序
			{"id": "asc"},      // 相同分数按ID排序
		},
	}

	// 对查询进行编码
	//var buf strings.Builder
	//if err := json.NewEncoder(&buf).Encode(query); err != nil {
	//	c.log.WithContext(ctx).Errorf("Failed to encode query: %v", err)
	//	return nil, 0, err
	//}
	queryBody, _ := json.Marshal(query)
	// 执行 Elasticsearch 查询
	//res, err := c.data.es.Search(
	//	c.data.es.Search.WithContext(ctx),
	//	c.data.es.Search.WithIndex("products"),
	//	c.data.es.Search.WithBody(strings.NewReader(buf.String())),
	//)
	res, err := esapi.SearchRequest{
		Index: []string{"products"},
		Body:  strings.NewReader(string(queryBody)),
	}.Do(ctx, c.data.es)
	if err != nil {
		c.log.Infof("Es商品搜索错误 = %+v", err)
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("ES返回错误: %s", res.String())
	}

	// 解析 Elasticsearch 查询结果
	var esResponse struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source struct {
					ID int `json:"column:id"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, 0, fmt.Errorf("解析ES响应失败: %v", err)
	}
	// 获取商品ID列表
	var ids []int
	for _, hit := range esResponse.Hits.Hits {
		ids = append(ids, hit.Source.ID)
	}
	total := esResponse.Hits.Total.Value
	if len(ids) == 0 {
		return []*biz.Content{}, total, nil
	}

	// 2. 查询MySQL获取完整数据
	var dbResults []*ContentDetail
	if err := c.data.db.Model(&ContentDetail{}).
		Where("id IN ?", ids).
		Find(&dbResults).Error; err != nil {
		c.log.Infof("Mysql商品搜索错误 = %+v", err)
		return nil, 0, err
	}
	// 按ES返回顺序排序
	idMap := make(map[int]int)
	for i, id := range ids {
		idMap[id] = i
	}
	sort.Slice(dbResults, func(i, j int) bool {
		return idMap[int(dbResults[i].ID)] < idMap[int(dbResults[j].ID)]
	})

	//构造返回请求
	var contents []*biz.Content
	//将数据库查找的结构，映射到biz.Content定义的结构
	for _, r := range dbResults {
		contents = append(contents, &biz.Content{
			ID:          int64(r.ID),
			Title:       r.Title,
			Description: r.Description,
			Picture_url: r.Picture_url,
			Price:       r.Price,
			Quantity:    r.Quantity,
			Categories:  strings.Split(r.Categories, ","), //字符串转数组
		})
	}

	return contents, total, nil
}

// 推荐商品：接入Gorse模型
func (c *contentRepo) Recommend(ctx context.Context, user_id int64, in_page, in_pageSize int32) ([]*biz.Content, int64, error) {
	var contents []*ContentDetail
	var total int64

	// 设置默认页大小
	var page, pageSize = 1, 10
	if in_page > 0 {
		page = int(in_page)
	}
	if in_pageSize > 0 {
		pageSize = int(in_pageSize)
	}
	offset := (page - 1) * pageSize

	// 获取总数
	if err := c.data.db.Model(&ContentDetail{}).Count(&total).Error; err != nil {
		c.log.Infof("Recommend获取商品总数错误")
		return nil, 0, err
	}

	// 查询商品数据并按ID逆序排序
	if err := c.data.db.Model(&ContentDetail{}).
		Order("id DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&contents).Error; err != nil {
		c.log.Infof("Recommend获取商品信息错误")
		return nil, 0, err
	}

	// 将数据库查找的结构，映射到biz.Content定义的结构
	var bizContents []*biz.Content
	for _, r := range contents {
		bizContents = append(bizContents, &biz.Content{
			ID:          int64(r.ID),
			Title:       r.Title,
			Description: r.Description,
			Picture_url: r.Picture_url,
			Price:       r.Price,
			Quantity:    r.Quantity,
			Categories:  strings.Split(r.Categories, ","), // 字符串转数组
		})
	}

	return bizContents, total, nil
}

//执行db
