package users

import (
	"github.com/gin-gonic/gin"
)


// /profile
func UsersProfiles(router *gin.RouterGroup) {
	router.GET("/:username", GetProfile)
}


// /users
func UsersRegister(router *gin.RouterGroup) {
	router.GET("/", GetUsers)
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
}

// /auth
func UsersValidation(router *gin.RouterGroup) {
	router.POST("/validate", ValidateUserToken)
}