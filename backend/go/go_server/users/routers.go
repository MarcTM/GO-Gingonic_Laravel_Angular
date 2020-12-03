package users

import (
	"github.com/gin-gonic/gin"
)


func UsersRegister(router *gin.RouterGroup) {
	router.GET("/test", TestUsers)
	router.GET("/", GetUsers)
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
}