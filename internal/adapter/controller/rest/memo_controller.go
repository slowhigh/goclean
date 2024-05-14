package rest

import (
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

func (c MemoController) FindOneMemo(input memo_dto.FindOneMemoInput) (int, *memo_dto.FindOneMemoOutput) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return http.StatusInternalServerError, nil
	}
	if !isExist {
		return http.StatusOK, nil
	}

	memo, ok := c.memoUcase.FindOne(input.ID)
	if !ok {
		return http.StatusInternalServerError, nil
	}

	output := memo_dto.NewFindOneMemoOutput(*memo)
	return http.StatusOK, &output
}

func (c MemoController) FindAllMemo(input memo_dto.FindAllMemoInput) (int, *dto.Paginated[[]memo_dto.FindAllMemoOutput]) {
	const perPage = 100

	totalCount, ok := c.memoUcase.CountMemo(input.Start, input.End, input.Keyword)
	if !ok {
		return http.StatusInternalServerError, nil
	}

	memos, ok := c.memoUcase.FindAllMemo(input.Start, input.End, input.Keyword, perPage, *input.Page)
	if !ok {
		return http.StatusInternalServerError, nil
	}

	memoOutput := make([]memo_dto.FindAllMemoOutput, len(*memos))
	for i, memo := range *memos {
		memoOutput[i] = *memo_dto.NewFindAllMemoOutput(memo)
	}

	output := dto.NewPaginatedOutput(memoOutput, *input.Page, perPage, *totalCount)
	return http.StatusOK, &output
}

func (c MemoController) CreateMemo(input memo_dto.CreateMemoInput) (int, *memo_dto.CreateMemoOutput) {
	memo, ok := c.memoUcase.CreateMemo(input.NewMemo())
	if !ok {
		return http.StatusInternalServerError, nil
	}

	output := memo_dto.NewCreateMemoOutput(*memo)
	return http.StatusOK, &output
}

func (c MemoController) UpdateMemo(input memo_dto.UpdateMemoInput) (int, *memo_dto.UpdateMemoOutput) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return http.StatusInternalServerError, nil
	}
	if !isExist {
		return http.StatusOK, nil
	}

	memo, ok := c.memoUcase.UpdateMemo(input.ID, input.NewMemo())
	if !ok {
		return http.StatusInternalServerError, nil
	}

	output := memo_dto.NewUpdateMemoOutput(*memo)
	return http.StatusOK, &output
}

func (c MemoController) DeleteMemo(input memo_dto.DeleteMemoInput) (int, *memo_dto.DeleteMemoOutput) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return http.StatusInternalServerError, nil
	}
	if !isExist {
		return http.StatusOK, nil
	}

	memo, ok := c.memoUcase.DeleteMemo(input.ID)
	if !ok {
		return http.StatusInternalServerError, nil
	}

	output := memo_dto.NewDeleteMemoOutput(*memo)
	return http.StatusOK, &output
}
