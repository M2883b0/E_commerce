package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"payment_system/internal/biz"
	"time"
)

type PaymentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &PaymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type PaymentDetail struct {
	//ID             int64         `gorm:"column:id;primary_key"`  // 自增ID
	OrderId   string    `gorm:"column:order_id;primary_key"`
	Amount    string    `gorm:"column:amount"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

// Todo 改表名
func (p PaymentDetail) TableName() string {
	return "ps_payment_info.t_payment_details"
}

func (p *PaymentRepo) Create(ctx context.Context, payment biz.Payment) error {
	detail := PaymentDetail{
		OrderId: payment.OrderID,
		Amount:  payment.Amount,
		Status:  payment.Status,
	}

	db := p.data.db
	if err := db.Create(&detail).Error; err != nil {
		p.log.Errorf("payment create error = %v", err)
		return err
	}
	return nil
}

func (p *PaymentRepo) Update(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p *PaymentRepo) FindByID(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

//func (c *paymentRepo) Create(ctx context.Context, content *biz.Payment) error {
//	c.log.Infof("contentRepo Create context = %+v", content)
//	detail := PaymentDetail{}
//	db := c.data.db
//	if err := db.Create(&detail).Error; err != nil {
//		c.log.Errorf("content create error = %v", err)
//		return err
//	}
//	return nil
//}

//
//func (c *paymentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
//	db := c.data.db
//	detail := ContentDetail{
//
//	}
//	if err := db.Where("id = ?", id).
//		Updates(&detail).Error; err != nil {
//		c.log.WithContext(ctx).Errorf("content update error = %v", err)
//		return err
//	}
//	return nil
//}
//
//func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
//	db := c.data.db
//	var detail ContentDetail
//	err := db.Where("id = ?", id).First(&detail).Error
//	if err == gorm.ErrRecordNotFound {
//		return false, nil
//	}
//	if err != nil {
//		c.log.WithContext(ctx).Errorf("ContentDao isExist = [%v]", err)
//		return false, err
//	}
//	return true, nil
//}
//
//func (c *contentRepo) Delete(ctx context.Context, id int64) error {
//	db := c.data.db
//	// 删除索引信息
//	err := db.Where("id = ?", id).
//		Delete(&ContentDetail{}).Error
//	if err != nil {
//		c.log.WithContext(ctx).Errorf("content delete error = %v", err)
//		return err
//	}
//	return nil
//}
//
//func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) ([]*biz.Content, int64, error) {
//	// 构造查询条件
//	query := c.data.db.Model(&ContentDetail{})
//	if params.ID != 0 {
//		query = query.Where("id = ?", params.ID)
//	}
//	if params.Author != "" {
//		query = query.Where("author = ?", params.Author)
//	}
//	if params.Title != "" {
//		query = query.Where("title = ?", params.Title)
//	}
//	// 总数
//	var total int64
//	if err := query.Count(&total).Error; err != nil {
//		return nil, 0, err
//	}
//	//设置默认页大小
//	var page, pageSize = 1, 10
//	if params.Page > 0 {
//		page = int(params.Page)
//	}
//	if params.PageSize > 0 {
//		pageSize = int(params.PageSize)
//	}
//	offset := (page - 1) * pageSize
//	//进行数据库查找
//	var results []*ContentDetail
//	if err := query.Offset(offset).Limit(pageSize).
//		Find(&results).Error; err != nil {
//		c.log.WithContext(ctx).Errorf("content find error = %v", err)
//		return nil, 0, err
//	}
//	var contents []*biz.Content
//	//将数据库查找的结构，映射到biz.Content定义的结构
//	for _, r := range results {
//		contents = append(contents, &biz.Content{
//			ID:             r.ID,
//			Title:          r.Title,
//			VideoURL:       r.VideoURL,
//			Author:         r.Author,
//			Description:    r.Description,
//			Thumbnail:      r.Thumbnail,
//			Category:       r.Category,
//			Duration:       r.Duration,
//			Resolution:     r.Resolution,
//			FileSize:       r.FileSize,
//			Format:         r.Format,
//			Quality:        r.Quality,
//			ApprovalStatus: r.ApprovalStatus,
//			UpdatedAt:      r.UpdatedAt,
//			CreatedAt:      r.CreatedAt,
//		})
//	}
//	return contents, total, nil
//}
