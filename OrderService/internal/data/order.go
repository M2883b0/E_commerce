package data

import (
	"OrderService/internal/biz"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

type OrderInfo struct {
	gorm.Model
	UserID        uint64         `gorm:"user_id"`
	PhoneNumber   string         `gorm:"phone_number"`
	IsPaid        string         `gorm:"is_paid"`
	StreetAddress string         `gorm:"street_address"`
	City          string         `gorm:"city"`
	Country       string         `gorm:"country"`
	ZipCode       uint32         `gorm:"zip_code"`
	OrderItems    datatypes.JSON `gorm:"type:json"`
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *orderRepo) Create(ctx context.Context, order *biz.Order) error {
	c.log.Infof("OrderInfo Create order = %+v", order)
	itemJSON, err := json.Marshal(order.OrderItems)
	if err != nil {
		return errors.Unwrap(err)
	}
	detail := OrderInfo{
		UserID:        order.UserID,
		PhoneNumber:   order.PhoneNumber,
		UserCurrency:  order.UserCurrency,
		IsPaid:        order.IsPaid,
		StreetAddress: order.StreetAddress,
		City:          order.City,
		Country:       order.Country,
		ZipCode:       order.ZipCode,
		OrderItems:    datatypes.JSON(itemJSON),
	}

	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("order create error = %v", err)
		return err
	}
	order.OrderId = uint64(detail.ID)
	c.log.Infof("===================%+v", order.OrderId)
	return nil
}

func (c *orderRepo) Update(ctx context.Context, id uint64, order *biz.Order) error {
	itemsJSON, err := json.Marshal(order.OrderItems)
	if err != nil {
		return errors.Unwrap(err)
	}
	c.log.Infof("OrderInfo Update order = %+v", order)
	detail := OrderInfo{
		UserID:        order.UserID,
		PhoneNumber:   order.PhoneNumber,
		UserCurrency:  order.UserCurrency,
		IsPaid:        order.IsPaid,
		StreetAddress: order.StreetAddress,
		City:          order.City,
		Country:       order.Country,
		ZipCode:       order.ZipCode,
		OrderItems:    datatypes.JSON(itemsJSON),
	}
	db := c.data.db
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("order update error = %v", err)
		return err
	}
	return nil
}

func (c *orderRepo) IsExist(ctx context.Context, id uint64) (bool, error) {
	db := c.data.db
	var detail OrderInfo
	err := db.Where("id = ?", id).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("Order isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *orderRepo) Delete(ctx context.Context, id uint64) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&OrderInfo{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("order delete error = %v", err)
		return err
	}
	return nil
}

func (c *orderRepo) Find(ctx context.Context, params *biz.FindParams) ([]*biz.Order, int64, error) {
	query := c.data.db.Model(&OrderInfo{})
	// 构造查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.PhoneNumber != "" {
		query = query.Where("phone_number = ?", params.PhoneNumber)
	}
	if params.IsPaid != nil {
		query = query.Where("is_paid = ?", *params.IsPaid)
	}
	// 总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//设置默认页大小
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = int(params.Page)
	}
	if params.PageSize > 0 {
		pageSize = int(params.PageSize)
	}
	offset := (page - 1) * pageSize
	var results []*OrderInfo
	if err := query.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user find error = %v", err)
		return nil, 0, err
	}
	// 转换逻辑
	var orders []*biz.Order
	for _, r := range results {
		var items []*biz.OrderItem
		if err := json.Unmarshal(r.OrderItems, &items); err != nil {
			return nil, total, errors.Unwrap(err)
		}
		orders = append(orders, &biz.Order{
			OrderId:       uint64(r.ID),
			UserID:        r.UserID,
			UserCurrency:  r.UserCurrency,
			PhoneNumber:   r.PhoneNumber,
			IsPaid:        r.IsPaid,
			StreetAddress: r.StreetAddress,
			City:          r.City,
			Country:       r.Country,
			ZipCode:       r.ZipCode,
			OrderItems:    items,
		})
	}
	return orders, total, nil
}
