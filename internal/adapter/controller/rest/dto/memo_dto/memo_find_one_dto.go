package memo_dto

import (
	"time"

	"github.com/slowhigh/goclean/internal/entity"
)

type FindOneMemoReq struct {
	ID int64
}

type FindOneMemoRes struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
}

func NewFindOneMemoRes(memo entity.Memo) FindOneMemoRes {
	return FindOneMemoRes{
		ID:        int64(memo.ID),
		CreatedAt: memo.CreatedAt,
		UpdatedAt: memo.UpdatedAt,
		Writer:    memo.Writer,
		Content:   memo.Content,
	}
}
