package routers

import (
	"SatelitePacheco/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var user models.Users

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status"	:  "ERROR",
			"message"	: err.Error(),
		})
		return
	}
	
	context.SetCookie("jwt","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",1,"APP/SERVICE/JWT","localhost",true,true)
	
	context.JSON(http.StatusOK, gin.H{
		"status"	: "Authorized",
		"message"	: "token",
	})
}
