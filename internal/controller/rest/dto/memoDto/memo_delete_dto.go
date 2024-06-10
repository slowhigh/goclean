package memoDto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type DeleteMemoReq struct {
	ID int64
}

type DeleteMemoRes struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewDeleteMemoRes(memo entity.Memo) DeleteMemoRes {
	return DeleteMemoRes{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
