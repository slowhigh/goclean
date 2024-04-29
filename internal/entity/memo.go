package entity

import (
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Writer  string `gorm:"column:writer;type:varchar;not null" validate:"required"`
	Content string `gorm:"column:content;type:varchar;not null" validate:"required"`
}
