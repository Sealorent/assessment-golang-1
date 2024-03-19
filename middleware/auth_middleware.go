package middleware

import (
	"final_project/common"
	"final_project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	// get the token from the header
	authorizationValue := ctx.GetHeader("Authorization")

	// check if the token is empty with the "Bearer " prefix
	splittedValue := strings.Split(authorizationValue, "Bearer ")
	if len(splittedValue) <= 1 {
		var r common.Response = common.Response{
			Success: false,
			Message: "Token not found",
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}
	jwtToken := splittedValue[1]

	claims, err := utils.GetJWTClaims(jwtToken)
	if err != nil {
		var r common.Response = common.Response{
			Success: false,
			Message: "Invalid token",
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}
