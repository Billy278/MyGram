package socialmedia

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	modelMedia "github.com/Billy278/MyGram/module/model/socialmedia"
	token "github.com/Billy278/MyGram/module/model/token"
	srvMedia "github.com/Billy278/MyGram/module/service/socialmedia"
	cripto "github.com/Billy278/MyGram/pkg/cripto"
	middleware "github.com/Billy278/MyGram/pkg/middleware"
	response "github.com/Billy278/MyGram/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MediaCtrlImpl struct {
	SrvMedia srvMedia.SocialMediaSrv
	Validate *validator.Validate
}

func NewMediaCtrlImpl(srvMedia srvMedia.SocialMediaSrv, validate *validator.Validate) MediaCtrl {
	return &MediaCtrlImpl{
		SrvMedia: srvMedia,
		Validate: validate,
	}
}

// @BasePath /api/v1
// @Summary Insert Media
// @Schemes http
// @Description Insert Social Media
// @Accept json
// @security Bearer
// @Param media body modelMedia.SocialMediaCreate true "media payload"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/media [post]
func (m_ctrl *MediaCtrlImpl) CtlInsertMedia(ctx *gin.Context) {
	reqIn := modelMedia.SocialMediaCreate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := m_ctrl.Validate.Struct(reqIn)
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
	res, err := m_ctrl.SrvMedia.SrvInsertMedia(ctx, reqIn)
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
		Message: "success created Social Media",
		Data:    "success created With  " + res.Name,
	})

}

// @BasePath /api/v1
// @Summary Update Media
// @Schemes http
// @Description Update Social Media
// @Accept json
// @security Bearer
// @Param media body modelMedia.SocialMediaUpdate true "media payload"
// @Param id path int true "idMedia"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/media/{idMedia} [put]
func (m_ctrl *MediaCtrlImpl) CtlUpdateMedia(ctx *gin.Context) {
	reqIn := modelMedia.SocialMediaUpdate{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: response.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := m_ctrl.Validate.Struct(reqIn)
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
	//get Id Media
	idMedia, err := m_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	reqIn.Id = idMedia
	//get user_id from JWT
	id, _ := strconv.Atoi(accessClaim.UserID)
	reqIn.User_id = uint64(id)
	res, err := m_ctrl.SrvMedia.SrvUpdateMedia(ctx, reqIn)
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
		Message: "success Updated Social Media",
		Data:    "success Updated With  " + res.Name,
	})
}

// @BasePath /api/v1
// @Summary FindBy Id Media
// @Schemes http
// @Description FindBy Id  Social Media
// @Accept json
// @security Bearer
// @Param id path int true "idMedia"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=modelMedia.SocialMediaRes}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/media/{idMedia} [get]
func (m_ctrl *MediaCtrlImpl) CtlFindByIdMedia(ctx *gin.Context) {
	id, err := m_ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := m_ctrl.SrvMedia.SrvFindByIdMedia(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Find Social Media By id",
		Data:    res,
	})
}

// @BasePath /api/v1
// @Summary FindAll  Media
// @Schemes http
// @Description FindAll  Social Media
// @Accept json
// @security Bearer
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]modelMedia.SocialMediaRes}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/media [get]
func (m_ctrl *MediaCtrlImpl) CtlFindAllMedia(ctx *gin.Context) {
	res, err := m_ctrl.SrvMedia.SrvFindAllMedia(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "success Find All Social Media",
		Data:    res,
	})
}

// @BasePath /api/v1
// @Summary Delete  Media
// @Schemes http
// @Description Delete  Social Media
// @Accept json
// @Param id path int true "idMedia"
// @security Bearer
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/media/{idMedia} [delete]
func (m_ctrl *MediaCtrlImpl) CtlDeleteMedia(ctx *gin.Context) {
	idMedia, err := m_ctrl.getIdFromParam(ctx)
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
	res, err := m_ctrl.SrvMedia.SrvDeleteMedia(ctx, idMedia, user_id)
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
		Message: "success Deleted Social Media",
		Data:    "success Deleted With  Id" + strconv.Itoa(int(res.Id)),
	})

}
func (m_ctrl *MediaCtrlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
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
