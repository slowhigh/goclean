package rest

import (
	"math"
	"net/http"

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
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return nil, http.StatusInternalServerError
	} else if !isExist {
		return nil, http.StatusNotFound
	}

	memo, ok := c.memoUcase.FindOne(input.ID)
	if !ok {
		return nil, http.StatusInternalServerError
	}

	return memo_dto.NewFindOneMemoOutput(memo), http.StatusOK
}

func (c MemoController) FindAllMemo(input memo_dto.FindAllMemoInput) (output *dto.Paginated[[]memo_dto.FindAllMemoOutput], status int) {
	const perPage = 100

	totalCount, ok := c.memoUcase.CountMemo(input.Start, input.End, input.Keyword)
	if !ok {
		return nil, http.StatusInternalServerError
	}

	memos, ok := c.memoUcase.FindAllMemo(input.Start, input.End, input.Keyword, perPage, *input.Page)
	if !ok {
		return nil, http.StatusInternalServerError
	}

	memoOutput := make([]memo_dto.FindAllMemoOutput, len(memos))
	for i, memo := range memos {
		memoOutput[i] = *memo_dto.NewFindAllMemoOutput(memo)
	}

	return &dto.Paginated[[]memo_dto.FindAllMemoOutput]{
		Page:       *input.Page,
		PerPage:    perPage,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(perPage))),
		TotalCount: totalCount,
		Data:       memoOutput,
	}, http.StatusOK
}

func (c MemoController) CreateMemo(input memo_dto.CreateMemoInput) (output *memo_dto.CreateMemoOutput, status int) {
	memo, ok := c.memoUcase.CreateMemo(input.NewMemo())
	if !ok {
		return nil, http.StatusInternalServerError
	}

	return memo_dto.NewCreateMemoOutput(memo), http.StatusOK
}

func (c MemoController) UpdateMemo(input memo_dto.UpdateMemoInput) (output *memo_dto.UpdateMemoOutput, status int) {
	memo, ok := c.memoUcase.UpdateMemo(input.ID, input.NewMemo())
	if !ok {
		return nil, http.StatusInternalServerError
	}

	return memo_dto.NewUpdateMemoOutput(memo), http.StatusOK
}

func (c MemoController) DeleteMemo(input memo_dto.DeleteMemoInput) (output *memo_dto.DeleteMemoOutput, status int) {
	memo, ok := c.memoUcase.DeleteMemo(input.ID)
	if !ok {
		return nil, http.StatusInternalServerError
	}

	return memo_dto.NewDeleteMemoOutput(memo), http.StatusOK
}
