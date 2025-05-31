package handler

import (
	"net/http"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/Amierza/warasin-mobile/backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	IMasterHandler interface {
		// Get Province & City
		GetAllProvince(ctx *gin.Context)
		GetAllCity(ctx *gin.Context)

		// Psycholog
		GetDetailPsycholog(ctx *gin.Context)
	}

	MasterHandler struct {
		masterService service.IMasterService
	}
)

func NewMasterHandler(masterService service.IMasterService) *MasterHandler {
	return &MasterHandler{
		masterService: masterService,
	}
}

// Get Province & City
func (mh *MasterHandler) GetAllProvince(ctx *gin.Context) {
	result, err := mh.masterService.GetAllProvince(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_PROVINCE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_PROVINCE,
		Data:     result.Data,
	}

	ctx.JSON(http.StatusOK, res)
}
func (mh *MasterHandler) GetAllCity(ctx *gin.Context) {
	var payload dto.CityQueryRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := mh.masterService.GetAllCity(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_CITY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_CITY,
		Data:     result.Data,
	}

	ctx.JSON(http.StatusOK, res)
}

// Psycholog
func (mh *MasterHandler) GetDetailPsycholog(ctx *gin.Context) {
	idStr := ctx.Param("id")
	result, err := mh.masterService.GetDetailPsycholog(ctx, idStr)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_PSYCHOLOG, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_PSYCHOLOG, result)
	ctx.JSON(http.StatusOK, res)
}
