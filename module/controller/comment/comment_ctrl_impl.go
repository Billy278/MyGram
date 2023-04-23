package comment

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	modelComment "github.com/Billy278/MyGram/module/model/comment"
	token "github.com/Billy278/MyGram/module/model/token"
	srvComment "github.com/Billy278/MyGram/module/service/comment"
	cripto "github.com/Billy278/MyGram/pkg/cripto"
	middleware "github.com/Billy278/MyGram/pkg/middleware"
	response "github.com/Billy278/MyGram/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentCtrlImpl struct {
	SrvComment srvComment.CommentSrv
	Validate   *validator.Validate
}

func NewCommentCtrlImpl(srvComment srvComment.CommentSrv, validate *validator.Validate) CommentCtrl {
	return &CommentCtrlImpl{
		SrvComment: srvComment,
		Validate:   validate,
	}
}

func (c_ctrl *CommentCtrlImpl) CtrlInsertComment(ctx *gin.Context) {
	reqIn := modelComment.CommentCreate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := c_ctrl.Validate.Struct(reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InvalidPayload,
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := cripto.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InternalServer,
		})
		return
	}

	//get user_id from JWT
	id, _ := strconv.Atoi(accessClaim.UserID)
	reqIn.User_id = uint64(id)
	res, err := c_ctrl.SrvComment.SrvInsertComment(ctx, reqIn)
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
		Message: "success created Comment",
		Data:    "success created With  Message" + res.Message,
	})
}
func (c_ctrl *CommentCtrlImpl) CtrlUpdateComment(ctx *gin.Context) {
	reqIn := modelComment.CommentUpdate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := c_ctrl.Validate.Struct(reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InvalidPayload,
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := cripto.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InternalServer,
		})
		return
	}
	//get Id Comment
	idComment, err := c_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	reqIn.Id = idComment
	//get user_id from JWT
	id, _ := strconv.Atoi(accessClaim.UserID)
	reqIn.User_id = uint64(id)
	res, err := c_ctrl.SrvComment.SrvUpdateComment(ctx, reqIn)
	if err != nil {
		if err.Error() == "UNAUTHORIZED" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: response.Unauthorized,
				Error:   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Updated Commnent",
		Data:    "success Updated With Message  " + res.Message,
	})
}
func (c_ctrl *CommentCtrlImpl) CtrlFindByIdComment(ctx *gin.Context) {
	id, err := c_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := c_ctrl.SrvComment.SrvFindByIdComment(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Comment By id",
		Data:    res,
	})
}
func (c_ctrl *CommentCtrlImpl) CtrlFindAllComment(ctx *gin.Context) {
	res, err := c_ctrl.SrvComment.SrvFindAllComment(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Find All Comment",
		Data:    res,
	})
}
func (c_ctrl *CommentCtrlImpl) CtrlDeleteComment(ctx *gin.Context) {
	idComment, err := c_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		log.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InvalidPayload,
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := cripto.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		log.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
			Error:   response.InternalServer,
		})
		return
	}

	//get user_id from JWT
	id, _ := strconv.Atoi(accessClaim.UserID)
	user_id := uint64(id)

	res, err := c_ctrl.SrvComment.SrvDeleteComment(ctx, idComment, user_id)
	if err != nil {
		if err.Error() == "UNAUTHORIZED" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: response.Unauthorized,
				Error:   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Deleted Comment",
		Data:    "success Deleted With Id" + strconv.Itoa(int(res.Id)),
	})
}
func (c_ctrl *CommentCtrlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	return idUint, err

}
