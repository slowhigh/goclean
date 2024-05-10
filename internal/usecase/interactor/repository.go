package interactor

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type MemoRepo interface {
	Exist(id int64) (bool, error)
	FindOne(id int64) (entity.Memo, error)
	Count(start *time.Time, end *time.Time, keyword *string) (int64, error)
	FindAll(start *time.Time, end *time.Time, keyword *string, perPage int, page int) ([]entity.Memo, error)
	Create(newMemo entity.Memo) (entity.Memo, error)
	Update(id int64, newMemo entity.Memo) (entity.Memo, error)
	Delete(id int64) (entity.Memo, error)
}
