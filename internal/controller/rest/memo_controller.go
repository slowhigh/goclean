package rest

import (
	"github.com/slowhigh/goclean/internal/controller/rest/dto"
	"github.com/slowhigh/goclean/internal/controller/rest/dto/memo_dto"
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

func (c MemoController) FindOneMemo(input memo_dto.FindOneMemoReq) (*memo_dto.FindOneMemoRes, bool) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return nil, false
	}
	if !isExist {
		return nil, true
	}

	memo, ok := c.memoUcase.FindOne(input.ID)
	if !ok {
		return nil, false
	}

	output := memo_dto.NewFindOneMemoRes(*memo)
	return &output, true
}

func (c MemoController) FindAllMemo(input memo_dto.FindAllMemoReq) (dto.Paginated[[]memo_dto.FindAllMemoRes], bool) {
	const perPage = 100
	data := dto.Paginated[[]memo_dto.FindAllMemoRes]{}

	totalCount, ok := c.memoUcase.CountMemo(input.Start, input.End, input.Keyword)
	if !ok {
		return data, false
	}

	memos, ok := c.memoUcase.FindAllMemo(input.Start, input.End, input.Keyword, perPage, *input.Page)
	if !ok {
		return data, false
	}

	memoOutput := make([]memo_dto.FindAllMemoRes, len(*memos))
	for i, memo := range *memos {
		memoOutput[i] = memo_dto.NewFindAllMemoRes(memo)
	}

	data = dto.NewPaginatedRes(memoOutput, *input.Page, perPage, *totalCount)
	return data, true
}

func (c MemoController) CreateMemo(input memo_dto.CreateMemoReq) (memo_dto.CreateMemoRes, bool) {
	memo, ok := c.memoUcase.CreateMemo(input.NewMemo())
	if !ok {
		return memo_dto.CreateMemoRes{}, false
	}

	output := memo_dto.NewCreateMemoRes(*memo)
	return output, true
}

func (c MemoController) UpdateMemo(input memo_dto.UpdateMemoReq) (*memo_dto.UpdateMemoRes, bool) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return nil, false
	}
	if !isExist {
		return nil, true
	}

	memo, ok := c.memoUcase.UpdateMemo(input.ID, input.NewMemo())
	if !ok {
		return nil, false
	}

	output := memo_dto.NewUpdateMemoRes(*memo)
	return &output, true
}

func (c MemoController) DeleteMemo(input memo_dto.DeleteMemoReq) (*memo_dto.DeleteMemoRes, bool) {
	isExist, ok := c.memoUcase.Exist(input.ID)
	if !ok {
		return nil, false
	}
	if !isExist {
		return nil, true
	}

	memo, ok := c.memoUcase.DeleteMemo(input.ID)
	if !ok {
		return nil, false
	}

	output := memo_dto.NewDeleteMemoRes(*memo)
	return &output, true
}
