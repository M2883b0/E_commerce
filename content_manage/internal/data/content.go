package data

import (
	"content_manage/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strings"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	gorm.Model
	Title       string `json:"column:title"`
	Description string `json:"column:description"`
	Picture_url string `json:"column:picture_url"`
	Price       uint32 `json:"column:price"`
	Quantity    uint32 `json:"column:quantity"`
	Categories  string `json:"column:categories"`
}

func (c ContentDetail) TableName() string {
	return "goods.product" //数据库的表名
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	c.log.Infof("contentRepo Create context = %+v", content)
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
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("content create error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
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
		c.log.WithContext(ctx).Errorf("content update error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("ContentDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&ContentDetail{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("content delete error = %v", err)
		return err
	}
	return nil
}

func (c *contentRepo) Get(ctx context.Context, id int64) (*biz.Content, error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("content get error = %v", err)
		return nil, err
	}
	content := &biz.Content{
		ID:          int64(detail.ID),
		Title:       detail.Title,
		Description: detail.Description,
		Picture_url: detail.Picture_url,
		Price:       detail.Price,
		Quantity:    detail.Quantity,
		Categories:  strings.Split(detail.Categories, ","),
	}
	return content, nil
}

// 搜索框，搜索商品，接入ElasticSearch + Kibana（查找）        同步更新，使用消息队列Canal监控mysql的内容的变更
func (c *contentRepo) Find(ctx context.Context, search string, in_page, in_pageSize int32) ([]*biz.Content, int64, error) {
	// 构造查询条件
	query := c.data.db.Model(&ContentDetail{})
	query = query.Where("title = %?%", search)

	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//设置默认页大小
	var page, pageSize = 1, 10
	if in_page > 0 {
		page = int(in_page)
	}
	if in_pageSize > 0 {
		pageSize = int(in_pageSize)
	}
	offset := (page - 1) * pageSize
	//进行数据库查找
	var results []*ContentDetail
	if err := query.Offset(offset).Limit(pageSize).
		Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("content find error = %v", err)
		return nil, 0, err
	}
	var contents []*biz.Content
	//将数据库查找的结构，映射到biz.Content定义的结构
	for _, r := range results {
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

// 推荐商品：接入Grose模型
func (c *contentRepo) Recommend(ctx context.Context, user_id int64, in_page, in_pageSize int32) ([]*biz.Content, int64, error) {
	return nil, 0, nil
}

//执行db
