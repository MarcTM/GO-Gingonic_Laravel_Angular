package recipes

import (
	"github.com/gin-gonic/gin"
)

// /recipes
func RecipesRoutes(router *gin.RouterGroup) {
	router.GET("/", GetRecipes)
	router.POST("/", IsAuthenticated(CreateRecipe))
	router.GET("/recipe/:id", GetRecipeByID)
	router.PUT("/recipe/:id", IsAuthenticated(UpdateRecipe))
	router.DELETE("/recipe/:id", IsAuthenticated(DeleteRecipe))

	router.POST("/favorited/:id", IsAuthenticated(IsFavorited))
	router.PUT("/favorite/:id", IsAuthenticated(FavoriteRecipe))
	router.PUT("/unfavorite/:id", IsAuthenticated(UnfavoriteRecipe))

	router.POST("/owns/:id", IsAuthenticated(OwnsRecipe))
}

// /comments
func CommentsRoutes(router *gin.RouterGroup) {
	// (:id refers to the recipe id)
	router.GET("/:id", GetComments)
	router.POST("/:id", IsAuthenticated(CreateComment))
}