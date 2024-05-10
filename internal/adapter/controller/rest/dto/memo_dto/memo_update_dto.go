package memo_dto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type UpdateMemoInput struct {
	ID      int64
	Writer  string `json:"writer" binding:"required,min=1,max=10"`
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

func (input UpdateMemoInput) NewMemo() entity.Memo {
	return entity.Memo{
		Writer:  input.Writer,
		Content: input.Content,
	}
}

type UpdateMemoOutput struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewUpdateMemoOutput(memo entity.Memo) *UpdateMemoOutput {
	return &UpdateMemoOutput{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
