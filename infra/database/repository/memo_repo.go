package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/slowhigh/goclean/infra/database"
	"github.com/slowhigh/goclean/internal/entity"
	"gorm.io/gorm"
)

type MemoRepo struct {
	db *database.Database
}

func NewMemo(db *database.Database) entity.MemoRepo {
	db.AutoMigrate(&entity.Memo{})

	return &MemoRepo{
		db: db,
	}
}

// Exist implements entity.MemoRepo.
func (clr *MemoRepo) Exist(id int64) (bool, error) {
	err := clr.db.Select("id").Take(&entity.Memo{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// FindOne implements entity.MemoRepo.
func (clr *MemoRepo) FindOne(id int64) (entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.Take(&memo, id).Error
	if err != nil {
		return entity.Memo{}, err
	}

	return memo, nil
}

// Count implements entity.MemoRepo.
func (clr *MemoRepo) Count(start, end *time.Time, keyword *string) (int64, error) {
	var (
		count      int64
		conditions = strings.Builder{}
		param      = make([]interface{}, 0)
	)

	if start != nil {
		conditions.WriteString(" AND created_at >= ?")
		param = append(param, *start)
	}

	if end != nil {
		conditions.WriteString(" AND created_at <= ?")
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
		return -1, err
	}

	return count, nil
}

// FindAll implements entity.MemoRepo.
func (clr *MemoRepo) FindAll(start *time.Time, end *time.Time, keyword *string, perPage int, page int) ([]entity.Memo, error) {
	var (
		memos      []entity.Memo
		conditions = strings.Builder{}
		param      = make([]interface{}, 0)
	)

	if start != nil {
		conditions.WriteString(" AND created_at >= ?")
		param = append(param, *start)
	}

	if end != nil {
		conditions.WriteString(" AND created_at <= ?")
		param = append(param, *end)
	}

	if keyword != nil {
		conditions.WriteString(" AND content LIKE ?")
		param = append(param, fmt.Sprintf("%%%s%%", *keyword))
	}

	err := clr.db.Model(&entity.Memo{}).
		Where(fmt.Sprintf("1=1%s", conditions.String()), param...).
		Order("created_at ASC").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Debug().
		Find(&memos).Error

	return memos, err
}

// Create implements entity.MemoRepo.
func (clr *MemoRepo) Create(newMemo entity.Memo) (entity.Memo, error) {
	err := clr.db.Create(&newMemo).Error
	if err != nil {
		return entity.Memo{}, err
	}

	return newMemo, nil
}

// Update implements entity.MemoRepo.
func (clr *MemoRepo) Update(id int64, newMemo entity.Memo) (entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.First(&memo, id).Error
	if err != nil {
		return entity.Memo{}, err
	}

	memo = newMemo
	memo.ID = uint(id)
	err = clr.db.Save(&memo).Error
	if err != nil {
		return entity.Memo{}, err
	}

	return memo, nil
}

// Delete implements entity.MemoRepo.
func (clr *MemoRepo) Delete(id int64) (entity.Memo, error) {
	var (
		memo entity.Memo
	)

	err := clr.db.First(&memo, id).Error
	if err != nil {
		return entity.Memo{}, err
	}

	err = clr.db.Delete(&entity.Memo{}, id).Error
	if err != nil {
		return entity.Memo{}, err
	}

	return memo, nil
}
