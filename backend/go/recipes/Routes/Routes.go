package Routes

import (
	"go_server/recipes"
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
		// Recipes
		recipes.RecipesRoutes(grp1.Group("/recipes"))

		// Comments
		recipes.CommentsRoutes(grp1.Group("/comments"))
	}
	return r
}
