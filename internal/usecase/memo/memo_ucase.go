package memo

import (
	"log/slog"
	"time"

	"github.com/slowhigh/goclean/internal/entity"
	"github.com/slowhigh/goclean/internal/usecase/interactor"
)

type MemoUsecase struct {
	memoRepo interactor.MemoRepo
}

func NewMemoUsecase(memoRepo interactor.MemoRepo) MemoUsecase {
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

func (mu MemoUsecase) FindOne(id int64) (memo entity.Memo, ok bool) {
	memo, err := mu.memoRepo.FindOne(id)
	if err != nil {
		slog.Error(err.Error())
		return entity.Memo{}, false
	}

	return memo, true
}

func (mu MemoUsecase) CountMemo(start, end *time.Time, keyword *string) (count int64, ok bool) {
	count, err := mu.memoRepo.Count(start, end, keyword)
	if err != nil {
		slog.Error(err.Error())
		return -1, false
	}

	return count, true
}

func (mu MemoUsecase) FindAllMemo(start, end *time.Time, keyword *string, perPage, page int) (memos []entity.Memo, ok bool) {
	memos, err := mu.memoRepo.FindAll(start, end, keyword, perPage, page)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return memos, true
}

func (mu MemoUsecase) CreateMemo(newMemo entity.Memo) (memo entity.Memo, ok bool) {
	memo, err := mu.memoRepo.Create(newMemo)
	if err != nil {
		slog.Error(err.Error())
		return entity.Memo{}, false
	}

	return memo, true
}

func (mu MemoUsecase) UpdateMemo(id int64, newMemo entity.Memo) (memo entity.Memo, ok bool) {
	memo, err := mu.memoRepo.Update(id, newMemo)
	if err != nil {
		slog.Error(err.Error())
		return entity.Memo{}, false
	}

	return memo, true
}

func (mu MemoUsecase) DeleteMemo(id int64) (memo entity.Memo, ok bool) {
	memo, err := mu.memoRepo.Delete(id)
	if err != nil {
		slog.Error(err.Error())
		return entity.Memo{}, false
	}

	return memo, true
}
