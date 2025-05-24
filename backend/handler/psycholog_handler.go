package handler

import (
	"net/http"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/Amierza/warasin-mobile/backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	IPsychologHandler interface {
		// Authentication
		Login(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)

		// Psycholog
		GetDetailPsycholog(ctx *gin.Context)
	}

	PsychologHandler struct {
		psychologService service.IPsychologService
		masterRepo       service.IMasterService
	}
)

func NewPsychologHandler(psychologService service.IPsychologService, masterRepo service.IMasterService) *PsychologHandler {
	return &PsychologHandler{
		psychologService: psychologService,
		masterRepo:       masterRepo,
	}
}

// Authentication
func (ph *PsychologHandler) Login(ctx *gin.Context) {
	var payload dto.PsychologLoginRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ph.psychologService.Login(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN_PSYCHOLOG, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN_PSYCHOLOG, result)
	ctx.JSON(http.StatusOK, res)
}
func (ph *PsychologHandler) RefreshToken(ctx *gin.Context) {
	var payload dto.RefreshTokenRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ph.psychologService.RefreshToken(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REFRESH_TOKEN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REFRESH_TOKEN, result)
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

// Psycholog
func (ph *PsychologHandler) GetDetailPsycholog(ctx *gin.Context) {
	result, err := ph.psychologService.GetDetailPsycholog(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_PSYCHOLOG, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_PSYCHOLOG, result)
	ctx.JSON(http.StatusOK, res)
}
