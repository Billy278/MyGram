package photo

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	modelPhoto "github.com/Billy278/MyGram/module/model/photo"
	token "github.com/Billy278/MyGram/module/model/token"
	srvPhoto "github.com/Billy278/MyGram/module/service/photo"
	cripto "github.com/Billy278/MyGram/pkg/cripto"
	middleware "github.com/Billy278/MyGram/pkg/middleware"
	response "github.com/Billy278/MyGram/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoCtrlImpl struct {
	SrvPhoto srvPhoto.PhotoSrv
	Validate *validator.Validate
}

func NewPhotoCtrlImpl(srvPhoto srvPhoto.PhotoSrv, validate *validator.Validate) PhotoCtrl {
	return &PhotoCtrlImpl{
		SrvPhoto: srvPhoto,
		Validate: validate,
	}
}

// @BasePath /api/v1
// @Summary Insert Photo
// @Schemes http
// @Description Insert Photo
// @Accept json
// @security Bearer
// @Param photo body modelPhoto.PhotoCreate true "photo payload"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/photo [post]
func (p_ctrl *PhotoCtrlImpl) CtrlInsertPhoto(ctx *gin.Context) {
	reqIn := modelPhoto.PhotoCreate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := p_ctrl.Validate.Struct(reqIn)
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
	reqIn.User_id = int64(id)
	res, err := p_ctrl.SrvPhoto.SrvInsertPhoto(ctx, reqIn)
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
		Data:    "success created With title " + res.Title,
	})
}

// @BasePath /api/v1
// @Summary Update Photo
// @Schemes http
// @Description Update Photo
// @Accept json
// @security Bearer
// @Param photo body modelPhoto.PhotoUpdate true "photo payload"
// @Param id path int true "idphoto"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/photo/{idphoto} [put]
func (p_ctrl *PhotoCtrlImpl) CtrlUpdatePhoto(ctx *gin.Context) {
	reqIn := modelPhoto.PhotoUpdate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := p_ctrl.Validate.Struct(reqIn)
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
	//get id
	idphoto, err := p_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	reqIn.Id = idphoto

	//get user_id from JWT
	id, _ := strconv.Atoi(accessClaim.UserID)
	reqIn.User_id = int64(id)

	res, err := p_ctrl.SrvPhoto.SrvUpdatePhoto(ctx, reqIn)
	if err != nil {
		if err.Error() == "UNAUTHORIZED" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: response.Unauthorized,
				Error:   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: response.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Update",
		Data:    "success Update With title " + res.Title,
	})

}

// @BasePath /api/v1
// @Summary Find By Id Photo
// @Schemes http
// @Description Find By Id Photo
// @Accept json
// @security Bearer
// @Param id path int true "idphoto"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=modelPhoto.PhotoRes}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/photo/{idphoto} [get]
func (p_ctrl *PhotoCtrlImpl) CtrlFindByIdPhoto(ctx *gin.Context) {
	id, err := p_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := p_ctrl.SrvPhoto.SrvFindByIdPhoto(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Find photo By id",
		Data:    res,
	})
}

// @BasePath /api/v1
// @Summary Find By All Photo
// @Schemes http
// @Description Find All Photo
// @Accept json
// @security Bearer
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]modelPhoto.PhotoRes}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/photo [get]
func (p_ctrl *PhotoCtrlImpl) CtrlFindAllPhoto(ctx *gin.Context) {
	res, err := p_ctrl.SrvPhoto.SrvFindAllPhoto(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Find All photo",
		Data:    res,
	})
}

// @BasePath /api/v1
// @Summary delete Photo
// @Schemes http
// @Description delete Photo
// @Accept json
// @security Bearer
// @Param id path int true "idphoto"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/photo/{idphoto} [delete]
func (p_ctrl *PhotoCtrlImpl) CtrlDeletePhoto(ctx *gin.Context) {
	idphoto, err := p_ctrl.getIdFromParam(ctx)
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
	user_id := int64(id)

	res, err := p_ctrl.SrvPhoto.SrvDeletePhoto(ctx, idphoto, uint64(user_id))
	if err != nil {
		if err.Error() == "UNAUTHORIZED" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: response.Unauthorized,
				Error:   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: response.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Delete",
		Data:    "success Delete With id " + strconv.Itoa(int(res.Id)),
	})
}
func (p_ctrl *PhotoCtrlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
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
