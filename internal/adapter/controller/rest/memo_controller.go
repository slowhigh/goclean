package rest

import (
	"math"

	"github.com/slowhigh/goclean/internal/adapter/controller/rest/dto"
	"github.com/slowhigh/goclean/internal/adapter/controller/rest/dto/memo_dto"
	"github.com/slowhigh/goclean/internal/usecase/memo"
)

type MemoController struct {
	memoUcase memo.MemoUsecase
}

func NewMemoController(msgUsecase memo.MemoUsecase) MemoController {
	return MemoController{
		memoUcase: msgUsecase,
	}
}

func (c MemoController) FindOneMemo(input memo_dto.FindOneMemoInput) (output *memo_dto.FindOneMemoOutput, status int) {
	memo, ok := c.memoUcase.FindOne(input.ID)
	if !ok {
		return nil, false
	}

	return memo_dto.NewFindOneMemoOutput(*memo), true
}

func (c MemoController) FindAllMemo(input memo_dto.FindAllMemoInput) (*dto.Paginated[[]memo_dto.FindAllMemoOutput], bool) {
	const perPage = 100

	totalCount, ok := c.memoUcase.CountMemo(input.Start, input.End, input.Keyword)
	if !ok {
		return nil, false
	}

	memos, ok := c.memoUcase.FindAllMemo(input.Start, input.End, input.Keyword, *input.Page, perPage)
	if !ok {
		return nil, false
	}

	memoOutput := make([]memo_dto.FindAllMemoOutput, len(*memos))
	for i, memo := range *memos {
		memoOutput[i] = *memo_dto.NewFindAllMemoOutput(memo)
	}

	return &dto.Paginated[[]memo_dto.FindAllMemoOutput]{
		Page:       *input.Page,
		PerPage:    perPage,
		TotalPage:  int(math.Ceil(float64(*totalCount) / float64(perPage))),
		TotalCount: *totalCount,
		Data:       memoOutput,
	}, true
}

func (c MemoController) CreateMemo(input memo_dto.CreateMemoInput) (*memo_dto.CreateMemoOutput, bool) {
	memo, ok := c.memoUcase.CreateMemo(input.NewMemo())
	if !ok {
		return nil, false
	}

	return memo_dto.NewCreateMemoOutput(*memo), true
}

func (c MemoController) UpdateMemo(input memo_dto.UpdateMemoInput) (*memo_dto.UpdateMemoOutput, bool) {
	memo, ok := c.memoUcase.UpdateMemo(*input.ID, input.NewMemo())
	if !ok {
		return nil, false
	}

	return memo_dto.NewUpdateMemoOutput(*memo), true
}

func (c MemoController) DeleteMemo(input memo_dto.DeleteMemoInput) (*memo_dto.DeleteMemoOutput, bool) {
	memo, ok := c.memoUcase.DeleteMemo(*input.ID)
	if !ok {
		return nil, false
	}

	return memo_dto.NewDeleteMemoOutput(*memo), true
}
