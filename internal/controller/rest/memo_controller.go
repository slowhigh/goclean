package rest

import (
	"github.com/slowhigh/goclean/internal/controller/rest/dto"
	"github.com/slowhigh/goclean/internal/controller/rest/dto/memoDto"
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

func (c MemoController) FindOneMemo(input memoDto.FindOneMemoReq) (*memoDto.FindOneMemoRes, bool) {
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

	output := memoDto.NewFindOneMemoRes(*memo)
	return &output, true
}

func (c MemoController) FindAllMemo(input memoDto.FindAllMemoReq) (dto.Paginated[[]memoDto.FindAllMemoRes], bool) {
	const perPage = 100
	data := dto.Paginated[[]memoDto.FindAllMemoRes]{}

	totalCount, ok := c.memoUcase.CountMemo(input.Start, input.End, input.Keyword)
	if !ok {
		return data, false
	}

	memos, ok := c.memoUcase.FindAllMemo(input.Start, input.End, input.Keyword, perPage, *input.Page)
	if !ok {
		return data, false
	}

	memoOutput := make([]memoDto.FindAllMemoRes, len(*memos))
	for i, memo := range *memos {
		memoOutput[i] = memoDto.NewFindAllMemoRes(memo)
	}

	data = dto.NewPaginatedRes(memoOutput, *input.Page, perPage, *totalCount)
	return data, true
}

func (c MemoController) CreateMemo(input memoDto.CreateMemoReq) (memoDto.CreateMemoRes, bool) {
	memo, ok := c.memoUcase.CreateMemo(input.NewMemo())
	if !ok {
		return memoDto.CreateMemoRes{}, false
	}

	output := memoDto.NewCreateMemoRes(*memo)
	return output, true
}

func (c MemoController) UpdateMemo(input memoDto.UpdateMemoReq) (*memoDto.UpdateMemoRes, bool) {
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

	output := memoDto.NewUpdateMemoRes(*memo)
	return &output, true
}

func (c MemoController) DeleteMemo(input memoDto.DeleteMemoReq) (*memoDto.DeleteMemoRes, bool) {
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

	output := memoDto.NewDeleteMemoRes(*memo)
	return &output, true
}
