package memo_dto

import "github.com/slowhigh/goclean/internal/entity"

type CreateMemoReq struct {
	Writer  string `json:"writer" binding:"required,min=1,max=10"`
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

func (input CreateMemoReq) NewMemo() entity.Memo {
	return entity.Memo{
		Writer:  input.Writer,
		Content: input.Content,
	}
}

type CreateMemoRes struct {
	Writer  string `json:"writer"`
	Content string `json:"content"`
}

func NewCreateMemoRes(memo entity.Memo) CreateMemoRes {
	return CreateMemoRes{
		Writer:  memo.Writer,
		Content: memo.Content,
	}
}
