package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"payment_system/internal/biz"
	"strconv"
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

func (p *PaymentRepo) Create(ctx context.Context, payment *biz.Payment) error {
	detail := PaymentDetail{
		OrderId: strconv.FormatInt(payment.OrderID, 10),
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

func (p *PaymentRepo) Update(ctx context.Context, payment *biz.Payment) error {
	log.Infof("更新订单表:%+v", payment)
	detail := PaymentDetail{
		OrderId: strconv.FormatInt(payment.OrderID, 10),
		Status:  payment.Status,
	}
	db := p.data.db
	if err := db.Where("order_id = ?", payment.OrderID).Updates(&detail).Error; err != nil {
		p.log.Errorf("payment Update error = %v", err)
		return err
	}
	panic("implement me")
}

func (p *PaymentRepo) FindByID(ctx context.Context, id int64) (*biz.Payment, error) {
	db := p.data.db
	var detail PaymentDetail
	// 根据 order_id 查找记录
	if err := db.Where("order_id = ?", strconv.FormatInt(id, 10)).First(&detail).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			p.log.Infof("记录未找到: order_id = %d", id)
			return nil, nil
		}
		p.log.Errorf("查找记录时出错: %v", err)
		return nil, err
	}

	// 将 PaymentDetail 转换为 biz.Payment
	payment := &biz.Payment{
		OrderID: id,
		Amount:  detail.Amount,
		Status:  detail.Status,
	}

	return payment, nil
}
