package memo_dto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type FindOneMemoInput struct {
	ID int64
}

type FindOneMemoOutput struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewFindOneMemoOutput(memo entity.Memo) FindOneMemoOutput {
	return FindOneMemoOutput{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
