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

		// Practice
		GetAllPractice(ctx *gin.Context)

		// Available Slot
		GetAllAvailableSlot(ctx *gin.Context)

		// Consultation
		GetAllConsultation(ctx *gin.Context)
	}

	PsychologHandler struct {
		psychologService service.IPsychologService
		masterService    service.IMasterService
	}
)

func NewPsychologHandler(psychologService service.IPsychologService, masterService service.IMasterService) *PsychologHandler {
	return &PsychologHandler{
		psychologService: psychologService,
		masterService:    masterService,
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

// Practice
func (ph *PsychologHandler) GetAllPractice(ctx *gin.Context) {
	result, err := ph.psychologService.GetAllPractice(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_PRACTICE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_PRACTICE, result)
	ctx.JSON(http.StatusOK, res)
}

// Available Slot
func (ph *PsychologHandler) GetAllAvailableSlot(ctx *gin.Context) {
	result, err := ph.psychologService.GetAllAvailableSlot(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_AVAILABLE_SLOT, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_AVAILABLE_SLOT, result)
	ctx.JSON(http.StatusOK, res)
}

// Consultation
func (ph *PsychologHandler) GetAllConsultation(ctx *gin.Context) {
	var payload dto.PaginationRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ph.psychologService.GetAllConsultationWithPagination(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_CONSULTATION, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:   true,
		Messsage: dto.MESSAGE_SUCCESS_GET_LIST_CONSULTATION,
		Data:     result.Data,
		Meta:     result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}
