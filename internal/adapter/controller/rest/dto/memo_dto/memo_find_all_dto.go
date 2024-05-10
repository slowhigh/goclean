package memo_dto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type FindAllMemoInput struct {
	Start   *time.Time `form:"start"`
	End     *time.Time `form:"end"`
	Keyword *string    `form:"keyword"`
	Page    *int       `form:"page" binding:"required"`
}

type FindAllMemoOutput struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewFindAllMemoOutput(memo entity.Memo) *FindAllMemoOutput {
	return &FindAllMemoOutput{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
