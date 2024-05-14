package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slowhigh/goclean/internal/adapter/controller/rest"
	"github.com/slowhigh/goclean/internal/adapter/controller/rest/dto/memo_dto"
)

func GetMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(ctrl.FindOneMemo(memo_dto.FindOneMemoInput{ID: *id}))
}

func GetMemos(c *gin.Context, ctrl rest.MemoController) {
	input, ok := validateQuery[memo_dto.FindAllMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(ctrl.FindAllMemo(*input))
}

func PostMemo(c *gin.Context, ctrl rest.MemoController) {
	input, ok := validateBody[memo_dto.CreateMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(ctrl.CreateMemo(*input))
}

func PutMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	input, ok := validateBody[memo_dto.UpdateMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	input.ID = *id

	c.JSON(ctrl.UpdateMemo(*input))
}

func DeleteMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(ctrl.DeleteMemo(memo_dto.DeleteMemoInput{ID: *id}))
}

func validateBody[T any](c *gin.Context) (*T, bool) {
	var input T

	if err := c.ShouldBind(&input); err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &input, true
}

func validateQuery[T any](c *gin.Context) (*T, bool) {
	var input T

	if err := c.ShouldBindQuery(&input); err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &input, true
}

func validateInt64Param(c *gin.Context, key string) (*int64, bool) {
	param := c.Param(key)
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &id, true
}
