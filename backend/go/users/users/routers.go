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
	router.GET("/me", IsAuthenticated(Me))
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)

	router.POST("/following/:username", IsAuthenticated(IsFollowing))
	router.PUT("/follow/:id", IsAuthenticated(FollowUser))
	router.PUT("/unfollow/:id", IsAuthenticated(UnfollowUser))
}

// /auth
func UsersValidation(router *gin.RouterGroup) {
	router.POST("/validate", ValidateUserToken)
}