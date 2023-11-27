package gin_pagination

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Pagination[T any] struct {
	List      []T   `json:"list"`
	Total     int64 `json:"total"`
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	query     *gorm.DB
	ctx       *gin.Context
}

func computeTotalPage(total int64, pageSize int) int {
	totalPage := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPage++
	}
	return totalPage
}

func New[T any](model T, c *gin.Context) *Pagination[T] {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", options.PageSizeDefaultVal))
	return &Pagination[T]{
		List:     make([]T, 0),
		PageSize: pageSize,
		PageNum:  pageNum,
		ctx:      c,
		query:    options.DB.Model(model).Order("updated_at desc"),
	}
}

func (p *Pagination[T]) Eq(field string) *Pagination[T] {
	value := p.ctx.Query(field)
	if value == "" {
		return p
	}
	p.query.Where(field+" = ?", value)
	return p
}

func (p *Pagination[T]) Neq(field string) *Pagination[T] {
	value := p.ctx.Query(field)
	if value == "" {
		return p
	}
	p.query.Where(field+" <> ?", value)
	return p
}

func (p *Pagination[T]) Gt(field string) *Pagination[T] {
	value := p.ctx.Query(field)
	if value == "" {
		return p
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		return p
	}
	p.query.Where(field+" > ?", val)
	return p
}

func (p *Pagination[T]) Lt(field string) *Pagination[T] {
	value := p.ctx.Query(field)
	if value == "" {
		return p
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		return p
	}
	p.query.Where(field+" < ?", val)
	return p
}

func (p *Pagination[T]) Like(field string) *Pagination[T] {
	value := p.ctx.Query(field)
	if value == "" {
		return p
	}
	p.query.Where(field+" LIKE ?", "%"+value+"%")
	return p
}

func (p *Pagination[T]) CB(w func(*gorm.DB)) *Pagination[T] {
	w(p.query)
	return p
}

func (p *Pagination[T]) Query() error {
	if options == nil {
		return errors.New("please initialize paging through init function first")
	}
	p.query.Count(&p.Total)
	err := p.query.Offset((p.PageNum - 1) * p.PageSize).Limit(p.PageSize).Find(&p.List).Error
	if err != nil {
		return err
	}
	p.TotalPage = computeTotalPage(p.Total, p.PageSize)
	return nil
}
