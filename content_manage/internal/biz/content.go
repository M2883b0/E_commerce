package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Content struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Picture_url string   `json:"picture_url"`
	Price       float32  `json:"price"`
	Quantity    uint32   `json:"quantity"`
	Categories  []string `json:"categories"`
}

type QuantityDetail struct {
	ID       int64 `json:"id"`
	Is_add   bool  `json:"is_add"`
	Quantity int32 `json:"quantity"`
}

type ContentRepo interface {
	Create(context.Context, *Content) error
	Update(context.Context, int64, *Content) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, string, int32, int32) ([]*Content, int64, error)
	Get(context.Context, []int64) ([]*Content, error)
	Recommend(context.Context, int64, int32, int32) ([]*Content, int64, error)
	UpdateQuantity(context.Context, []*QuantityDetail) (bool, error)
}

// ContentUsecase is a Content usecase.
type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateContent creates a Content, and returns the new Content.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) error {
	return uc.repo.Create(ctx, c)
}

// UpdateContent update a Content.
func (uc *ContentUsecase) UpdateContent(ctx context.Context, c *Content) error {
	return uc.repo.Update(ctx, c.ID, c)
}

// DeleteContent delete a Content.
func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	ok, err := uc.repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("内容不存在,无法删除")
	}
	//内容存在，执行删除操作
	return uc.repo.Delete(ctx, id)
}

// FindContent find Content.
func (uc *ContentUsecase) FindContent(ctx context.Context, query string, page, pageSize int32) ([]*Content, int64, error) {
	repo := uc.repo
	contents, total, err := repo.Find(ctx, query, page, pageSize) //调用data层的Find实现
	if err != nil {
		return nil, 0, err
	}
	return contents, total, nil
}

func (uc *ContentUsecase) GetContent(ctx context.Context, ids []int64) ([]*Content, error) {
	repo := uc.repo
	contents, err := repo.Get(ctx, ids) //调用data层的Find实现
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (uc *ContentUsecase) RecommendContent(ctx context.Context, user_id int64, page int32, pageSize int32) ([]*Content, int64, error) {
	repo := uc.repo
	contents, total, err := repo.Recommend(ctx, user_id, page, pageSize) //调用data层的Find实现
	if err != nil {
		return nil, 0, err
	}
	return contents, total, nil
}

func (uc *ContentUsecase) UpdateContentQuantity(ctx context.Context, quantity_list []*QuantityDetail) (bool, error) {
	repo := uc.repo
	res, err := repo.UpdateQuantity(ctx, quantity_list) //调用data层的Find实现
	if err != nil {
		return false, err
	}
	return res, nil
}

//执行组合逻辑
