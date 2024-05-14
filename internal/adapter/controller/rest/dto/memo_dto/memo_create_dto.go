package memo_dto

import "github.com/slowhigh/goclean/internal/entity"

type CreateMemoInput struct {
	Writer  string `json:"writer" binding:"required,min=1,max=10"`
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

func (input CreateMemoInput) NewMemo() entity.Memo {
	return entity.Memo{
		Writer:  input.Writer,
		Content: input.Content,
	}
}

type CreateMemoOutput struct {
	Writer  string `json:"writer"`
	Content string `json:"content"`
}

func NewCreateMemoOutput(memo entity.Memo) CreateMemoOutput {
	return CreateMemoOutput{
		Writer:  memo.Writer,
		Content: memo.Content,
	}
}
