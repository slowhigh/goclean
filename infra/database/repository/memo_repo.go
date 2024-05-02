package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/slowhigh/goclean/infra/database"
	"github.com/slowhigh/goclean/internal/entity"
	"github.com/slowhigh/goclean/internal/usecase/interactor"
)

type MemoRepo struct {
	db *database.Database
}

func NewMemo(db *database.Database) interactor.MemoRepo {
	return &MemoRepo{
		db: db,
	}
}

// FindOne implements interactor.MemoRepo.
func (clr *MemoRepo) FindOne(id int64) (*entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.First(&memo, id).Error
	if err != nil {
		return nil, err
	}

	return &memo, nil
}

// Count implements interactor.MemoRepo.
func (clr *MemoRepo) Count(start, end *time.Time, keyword *string) (*int64, error) {
	var (
		count      int64
		conditions = strings.Builder{}
		param      = make([]interface{}, 0)
	)

	if start != nil {
		conditions.WriteString(" AND timestamp >= ?")
		param = append(param, *start)
	}

	if end != nil {
		conditions.WriteString(" AND timestamp <= ?")
		param = append(param, *end)
	}

	if keyword != nil {
		conditions.WriteString(" AND content LIKE ?")
		param = append(param, fmt.Sprintf("%%%s%%", *keyword))
	}

	err := clr.db.Model(&entity.Memo{}).
		Where(fmt.Sprintf("1=1%s", conditions.String()), param...).
		Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// FindAll implements interactor.MemoRepo.
func (clr *MemoRepo) FindAll(start *time.Time, end *time.Time, keyword *string, perPage int, page int) (*[]entity.Memo, error) {
	var (
		memos      []entity.Memo
		conditions = strings.Builder{}
		param      = make([]interface{}, 0)
	)

	if start != nil {
		conditions.WriteString(" AND timestamp >= ?")
		param = append(param, *start)
	}

	if end != nil {
		conditions.WriteString(" AND timestamp <= ?")
		param = append(param, *end)
	}

	if keyword != nil {
		conditions.WriteString(" AND content LIKE ?")
		param = append(param, fmt.Sprintf("%%%s%%", *keyword))
	}

	err := clr.db.Model(&entity.Memo{}).
		Where(fmt.Sprintf("1=1%s", conditions.String()), param...).
		Order("timestamp DESC").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Debug().
		Find(&memos).Error

	return &memos, err
}

// Create implements interactor.MemoRepo.
func (clr *MemoRepo) Create(newMemo entity.Memo) (*entity.Memo, error) {
	err := clr.db.Create(&newMemo).Error
	if err != nil {
		return nil, err
	}

	return &newMemo, nil
}

// Update implements interactor.MemoRepo.
func (clr *MemoRepo) Update(newMemo entity.Memo) (*entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.First(&memo, newMemo.ID).Error
	if err != nil {
		return nil, err
	}

	memo = newMemo
	err = clr.db.Save(&memo).Error
	if err != nil {
		return nil, err
	}

	return &memo, nil
}

// Delete implements interactor.MemoRepo.
func (clr *MemoRepo) Delete(id int64) (*entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.First(&memo, id).Error
	if err != nil {
		return nil, err
	}

	err = clr.db.Delete(&entity.Memo{}, id).Error
	if err != nil {
		return nil, err
	}

	return &memo, nil
}
