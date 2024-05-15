package entity

import (
	"time"

	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Writer  string `gorm:"column:writer;type:varchar;not null" validate:"required"`
	Content string `gorm:"column:content;type:varchar;not null" validate:"required"`
}

type MemoRepo interface {
	Exist(id int64) (bool, error)
	FindOne(id int64) (Memo, error)
	Count(start *time.Time, end *time.Time, keyword *string) (int64, error)
	FindAll(start *time.Time, end *time.Time, keyword *string, perPage int, page int) ([]Memo, error)
	Create(newMemo Memo) (Memo, error)
	Update(id int64, newMemo Memo) (Memo, error)
	Delete(id int64) (Memo, error)
}
