package middleware

import (
	"net/http"
	"strings"

	tokenModel "github.com/Billy278/MyGram/module/model/token"
	"github.com/Billy278/MyGram/pkg/cripto"
	response "github.com/Billy278/MyGram/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	HeaderKey  string
	ContextKey string
)

const (
	Authorization HeaderKey = "Authorization"

	AccessClaim ContextKey = "access_claim"

	BasicAuth  string = "Basic "
	BearerAuth string = "Bearer "
)

func BearerOAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth header
		header := ctx.GetHeader(string(Authorization))
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Token Not Found",
				Error:   response.Unauthorized,
			})
			return
		}
		//get token
		token := strings.Split(header, BearerAuth)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Token Not Found",
				Error:   response.Unauthorized,
			})
			return
		}

		// header token is found
		var claim tokenModel.AccessClaim
		err := cripto.ParseJWT(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
				Error:   response.SomethingWentWrong,
			})
			return
		}
		ctx.Set(string(AccessClaim), claim)
		ctx.Next()

	}

}
