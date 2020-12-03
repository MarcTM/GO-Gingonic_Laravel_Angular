package Routes

import (
	"go_server/users"
	"go_server/recipes"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		// Users
		users.UsersRegister(grp1.Group("/users"))

		// Recipes
		recipes.RecipesRegister(grp1.Group("/recipes"))
	}
	
	return r
}
