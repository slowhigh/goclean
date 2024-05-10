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
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error("param is not int type.")
		c.Status(http.StatusBadRequest)
		return
	}

	output, status := ctrl.FindOneMemo(memo_dto.FindOneMemoInput{ID: id})
	if output == nil {
		c.Status(status)
		return
	}

	c.JSON(status, *output)
}

func GetMemos(c *gin.Context, ctrl rest.MemoController) {
	input, ok := validateQuery[memo_dto.FindAllMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	output, status := ctrl.FindAllMemo(*input)
	if output == nil {
		c.Status(status)
		return
	}

	c.JSON(status, *output)
}

func PostMemo(c *gin.Context, ctrl rest.MemoController) {
	input, ok := validateBody[memo_dto.CreateMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	output, status := ctrl.CreateMemo(*input)
	if output == nil {
		c.Status(status)
		return
	}

	c.JSON(status, *output)
}

func PutMemo(c *gin.Context, ctrl rest.MemoController) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error("param is not int type.")
		c.Status(http.StatusBadRequest)
		return
	}

	input, ok := validateBody[memo_dto.UpdateMemoInput](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	input.ID = id
	output, status := ctrl.UpdateMemo(*input)
	if output == nil {
		c.Status(status)
		return
	}

	c.JSON(status, *output)
}
func DeleteMemo(c *gin.Context, ctrl rest.MemoController) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error("param is not int type.")
		c.Status(http.StatusBadRequest)
		return
	}

	output, status := ctrl.DeleteMemo(memo_dto.DeleteMemoInput{ID: id})
	if output == nil {
		c.Status(status)
		return
	}

	c.JSON(status, *output)
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