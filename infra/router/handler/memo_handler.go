package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slowhigh/goclean/internal/controller/rest/dto/memo_dto"
	"github.com/slowhigh/goclean/internal/controller/rest"
)

// GetMemo
//
// @Summary			Get memo
// @Description		get memo
// @Tags			memo
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		id path int64 true "ID search by id"
// @Success			200 {object} memo_dto.FindOneMemoRes
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/v1/memo/{id} [get]
func GetMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	res, ok := ctrl.FindOneMemo(memo_dto.FindOneMemoReq{ID: *id})
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListMemos
//
// @Summary			List memo
// @Description		list memos
// @Tags			memo
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		start		query	string	false	"createAt search by start"		example("2021-02-18T21:54:42.123Z")
// @Param   		end			query	string	false	"createAt search by end"		example("2021-02-18T21:54:42.123Z")
// @Param   		keyword		query	string	false	"content search by keyword"
// @Param   		page		query	int		true	"content search by keyword"		minimum(1)
// @Success			200 {object} memo_dto.FindAllMemoRes
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/v1/memo [get]
func ListMemos(c *gin.Context, ctrl rest.MemoController) {
	req, ok := validateQuery[memo_dto.FindAllMemoReq](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	res, ok := ctrl.FindAllMemo(*req)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateMemo
//
// @Summary			Create memo
// @Description		create memo
// @Tags			memo
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		request body memo_dto.CreateMemoReq true "writer and content for the new memo"
// @Success			200 {object} memo_dto.CreateMemoRes
// @Failure			400
// @Failure			500
// @Router			/v1/memo [post]
func CreateMemo(c *gin.Context, ctrl rest.MemoController) {
	req, ok := validateBody[memo_dto.CreateMemoReq](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	res, ok := ctrl.CreateMemo(*req)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateMemo
//
// @Summary			Update memo
// @Description		update memo
// @Tags			memo
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		id path int64 true "memo update by id" minimum(1)
// @Param   		req body memo_dto.UpdateMemoReq true "writer and content for the new memo"
// @Success			200 {object} memo_dto.UpdateMemoRes
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/v1/memo/{id} [put]
func UpdateMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	req, ok := validateBody[memo_dto.UpdateMemoReq](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}
	req.ID = *id

	res, ok := ctrl.UpdateMemo(*req)
	if !ok {
		c.Status(http.StatusInternalServerError)
	}
	if res == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteMemo
//
// @Summary			Delete memo
// @Description		delete memo
// @Tags			memo
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		id path int64 true "memo delete by id" minimum(1)
// @Success			200 {object} memo_dto.DeleteMemoRes
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/v1/memo/{id} [delete]
func DeleteMemo(c *gin.Context, ctrl rest.MemoController) {
	id, ok := validateInt64Param(c, "id")
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	res, ok := ctrl.DeleteMemo(memo_dto.DeleteMemoReq{ID: *id})
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}
	if res == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, res)
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
