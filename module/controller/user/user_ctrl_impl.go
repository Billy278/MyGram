package user

import (
	"net/http"

	modelUser "github.com/Billy278/MyGram/module/model/user"
	srvUser "github.com/Billy278/MyGram/module/service/user"
	response "github.com/Billy278/MyGram/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlImpl struct {
	SrvUser  srvUser.USerServ
	Validate *validator.Validate
}

func NewUserCtrlImpl(srvuser srvUser.USerServ, validate *validator.Validate) UserCtrl {
	return &UserCtrlImpl{
		SrvUser:  srvuser,
		Validate: validate,
	}
}

func (u_ctrl *UserCtrlImpl) Registration(ctx *gin.Context) {
	reqIn := modelUser.UserCreate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//validasi request
	err := u_ctrl.Validate.Struct(reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi username unix
	err = u_ctrl.SrvUser.SrvFindByUsername(ctx, reqIn.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi email unix
	err = u_ctrl.SrvUser.SrvFindByEmail(ctx, reqIn.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	res, err := u_ctrl.SrvUser.SrvInsertUser(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: response.InternalServer,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success created",
		Data:    "success created With Username " + res.Username,
	})

}
func (u_ctrl *UserCtrlImpl) LoginUser(ctx *gin.Context) {
	reqIn := modelUser.UserLogin{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//validasi request
	err := u_ctrl.Validate.Struct(reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	tokens, err := u_ctrl.SrvUser.SrvLoginUsername(ctx, reqIn.Username, reqIn.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: response.InternalServer,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success created",
		Data:    tokens,
	})
}
