package memo_dto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type DeleteMemoInput struct {
	ID int64
}

type DeleteMemoOutput struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewDeleteMemoOutput(memo entity.Memo) *DeleteMemoOutput {
	return &DeleteMemoOutput{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
