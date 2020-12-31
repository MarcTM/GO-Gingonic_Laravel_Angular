package recipes

import (
	"github.com/gin-gonic/gin"
)

// /recipes
func RecipesRoutes(router *gin.RouterGroup) {
	router.GET("/", GetRecipes)
	router.POST("/", IsAuthenticated(CreateRecipe))
	router.GET("/recipe/:id", GetRecipeByID)
	router.PUT("/recipe/:id", UpdateRecipe)
	router.DELETE("/recipe/:id", DeleteRecipe)

	router.POST("/favorited/:id", IsAuthenticated(IsFavorited))
	router.PUT("/favorite/:id", IsAuthenticated(FavoriteRecipe))
	router.PUT("/unfavorite/:id", IsAuthenticated(UnfavoriteRecipe))
}