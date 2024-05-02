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

func (mu MemoUsecase) FindOne(id int64) (*entity.Memo, bool) {
	memo, err := mu.memoRepo.FindOne(id)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return memo, true
}

func (mu MemoUsecase) CountMemo(start, end *time.Time, keyword *string) (*int64, bool) {
	count, err := mu.memoRepo.Count(start, end, keyword)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return count, true
}

func (mu MemoUsecase) FindAllMemo(start, end *time.Time, keyword *string, page int, perPage int) (*[]entity.Memo, bool) {
	memos, err := mu.memoRepo.FindAll(start, end, keyword, page, perPage)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return memos, true
}

func (mu MemoUsecase) CreateMemo(newMemo entity.Memo) *entity.Memo {
	memo, err := mu.memoRepo.Create(newMemo)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return memo
}

func (mu MemoUsecase) UpdateMemo(newMemo entity.Memo) *entity.Memo {
	memo, err := mu.memoRepo.Update(newMemo)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return memo
}

func (mu MemoUsecase) DeleteMemo(id int64) *entity.Memo {
	memo, err := mu.memoRepo.Delete(id)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return memo
}
