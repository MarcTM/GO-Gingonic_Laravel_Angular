package Routes

import (
	"go_server/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AddAllowHeaders("*")
	r.Use(cors.New(config))

	grp1 := r.Group("/api")
	{
		// Users
		users.UsersRegister(grp1.Group("/users"))

		// Profiles
		users.UsersProfiles(grp1.Group("/profiles"))

		// Auth
		users.UsersValidation(grp1.Group("/auth"))
	}
	return r
}
