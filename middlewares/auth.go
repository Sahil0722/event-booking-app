package middlewares

import (
	"net/http"

	"github.com/Sahil0722/event-booking-app/utils"
	"github.com/gin-gonic/gin"
)

func Authorize(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!!"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!!"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
