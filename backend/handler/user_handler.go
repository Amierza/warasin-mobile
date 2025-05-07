package handler

import (
	"net/http"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/service"
	"github.com/Amierza/warasin-mobile/backend/utils"
	"github.com/gin-gonic/gin"
)

type (
	IUserHandler interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
		SendForgotPasswordEmail(ctx *gin.Context)
		ForgotPassword(ctx *gin.Context)
		UpdatePassword(ctx *gin.Context)
		SendVerificationEmail(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
		GetDetailUser(ctx *gin.Context)
		UpdateUser(ctx *gin.Context)
		GetAllProvince(ctx *gin.Context)
	}

	UserHandler struct {
		userService service.IUserService
	}
)

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Register(ctx *gin.Context) {
	var payload dto.UserRegisterRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.Register(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) Login(ctx *gin.Context) {
	var payload dto.UserLoginRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.Login(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) RefreshToken(ctx *gin.Context) {
	var payload dto.RefreshTokenRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.RefreshToken(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REFRESH_TOKEN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REFRESH_TOKEN, result)
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

func (uh *UserHandler) SendForgotPasswordEmail(ctx *gin.Context) {
	var payload dto.SendForgotPasswordEmailRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := uh.userService.SendForgotPasswordEmail(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEND_FORGOT_PASSWORD_EMAIL, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SEND_FORGOT_PASSWORD_EMAIL, nil)
	ctx.JSON(http.StatusOK, res)
}

func (uh UserHandler) ForgotPassword(ctx *gin.Context) {
	var payload dto.ForgotPasswordRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.ForgotPassword(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CHECK_FORGOT_PASSWORD_TOKEN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CHECK_FORGOT_PASSWORD_TOKEN, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) UpdatePassword(ctx *gin.Context) {
	var payload dto.UpdatePasswordRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.UpdatePassword(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PASSWORD, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PASSWORD, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) SendVerificationEmail(ctx *gin.Context) {
	var payload dto.SendVerificationEmailRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := uh.userService.SendVerificationEmail(ctx.Request.Context(), payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEND_VERIFICATION_EMAIL, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SEND_VERIFICATION_EMAIL, nil)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) VerifyEmail(ctx *gin.Context) {
	var payload dto.VerifyEmailRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.VerifyEmail(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_EMAIL, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VERIFY_EMAIL, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) GetDetailUser(ctx *gin.Context) {
	result, err := uh.userService.GetDetailUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DETAIL_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DETAIL_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var payload dto.UpdateUserRequest
	if err := ctx.ShouldBind(&payload); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := uh.userService.UpdateUser(ctx, payload)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (ah *UserHandler) GetAllProvince(ctx *gin.Context) {
	result, err := ah.userService.GetAllProvince(ctx.Request.Context())
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
