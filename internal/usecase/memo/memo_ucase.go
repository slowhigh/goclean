package memo

import (
	"log/slog"
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type MemoUsecase struct {
	memoRepo entity.MemoRepo
}

func NewMemoUsecase(memoRepo entity.MemoRepo) MemoUsecase {
	return MemoUsecase{
		memoRepo: memoRepo,
	}
}

func (mu MemoUsecase) Exist(id int64) (isExist bool, ok bool) {
	isExist, err := mu.memoRepo.Exist(id)
	if err != nil {
		slog.Error(err.Error())
		return false, false
	}

	return isExist, true
}

func (mu MemoUsecase) FindOne(id int64) (*entity.Memo, bool) {
	memo, err := mu.memoRepo.FindOne(id)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &memo, true
}

func (mu MemoUsecase) CountMemo(start, end *time.Time, keyword *string) (*int64, bool) {
	count, err := mu.memoRepo.Count(start, end, keyword)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &count, true
}

func (mu MemoUsecase) FindAllMemo(start, end *time.Time, keyword *string, perPage, page int) (*[]entity.Memo, bool) {
	memos, err := mu.memoRepo.FindAll(start, end, keyword, perPage, page)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &memos, true
}

func (mu MemoUsecase) CreateMemo(newMemo entity.Memo) (*entity.Memo, bool) {
	memo, err := mu.memoRepo.Create(newMemo)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &memo, true
}

func (mu MemoUsecase) UpdateMemo(id int64, newMemo entity.Memo) (*entity.Memo, bool) {
	memo, err := mu.memoRepo.Update(id, newMemo)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &memo, true
}

func (mu MemoUsecase) DeleteMemo(id int64) (*entity.Memo, bool) {
	memo, err := mu.memoRepo.Delete(id)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &memo, true
}
