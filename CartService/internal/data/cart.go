package data

import (
	"CartService/api/operate"
	"CartService/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

type CartInfo struct {
	UserID    uint64 `gorm:"user_id"`
	ProductID uint64 `gorm:"product_id"`
	Quantity  uint64 `gorm:"quantity"`
}

// NewCartRepo NewOrderRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cartRepo) Create(ctx context.Context, cartItem *biz.CartItem) error {
	c.log.Infof("CartInfo Create order = %+v", cartItem)
	detail := CartInfo{
		UserID:    cartItem.UserID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
	}

	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("order create error = %v", err)
		return err
	}
	return nil
}

func (c *cartRepo) Update(ctx context.Context, cartItem *biz.CartItem) error {
	c.log.Infof("CartInfo Update order = %+v", cartItem)
	detail := CartInfo{
		Quantity: cartItem.Quantity,
	}
	db := c.data.db
	if err := db.Where("user_id = ? and product_id = ?", cartItem.UserID, cartItem.ProductID).Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("order update error = %v", err)
		return err
	}
	return nil
}

func (c *cartRepo) IsExist(ctx context.Context, cartItem *biz.CartItem) bool {
	db := c.data.db
	var detail CartInfo
	err := db.Where("user_id = ? and product_id = ?", cartItem.UserID, cartItem.ProductID).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	if err != nil {
		c.log.WithContext(ctx).Infof("Order isExist = [%v]", err)
		return false
	}
	return true
}

func (c *cartRepo) Delete(ctx context.Context, cartItem *biz.CartItem) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("user_id = ? and product_id = ?", cartItem.UserID, cartItem.ProductID).
		Delete(&CartInfo{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("order delete error = %v", err)
		return err
	}
	return nil
}

func (c *cartRepo) FindCartByUserId(ctx context.Context, params *biz.FindParams) ([]*biz.CartItem, int64, error) {
	query := c.data.db.Model(&CartInfo{})
	// 构造查询条件
	if params.ProductId == 0 {
		query = query.Where("user_id = ?", params.UserId)
	} else {
		query = query.Where("user_id = ? and product_id = ?", params.UserId, params.ProductId)
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
	var results []*CartInfo
	if err := query.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user find error = %v", err)
		return nil, 0, err
	}
	// 转换逻辑
	var cartItems []*biz.CartItem
	for _, r := range results {
		cartItems = append(cartItems, &biz.CartItem{
			UserID:    r.UserID,
			ProductID: r.ProductID,
			Quantity:  r.Quantity,
		})
	}
	return cartItems, total, nil
}

func (c *cartRepo) GetContentInfoById(ctx context.Context, id uint64) (*biz.ContentInfo, error) {
	response, err := c.data.contentClient.GetContent(ctx, &operate.GetContentReq{Id: int64(id)})
	if err != nil {
		return nil, err
	}
	var contentInfo = biz.ContentInfo{
		Id:             response.Contents.Id,
		Title:          response.Contents.Title,
		Description:    response.Contents.Description,
		PictureUrl:     response.Contents.PictureUrl,
		Price:          response.Contents.Price,
		ServerQuantity: response.Contents.Quantity,
		Categories:     response.Contents.GetCategories(),
	}

	return &contentInfo, nil
}
