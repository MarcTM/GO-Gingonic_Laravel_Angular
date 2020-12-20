package recipes

import (
	"github.com/gin-gonic/gin"
)

// /recipes
func RecipesRegister(router *gin.RouterGroup) {
	router.GET("/", GetRecipes)
	router.POST("/", IsAuthenticated(CreateRecipe))
	router.GET("/:id", GetRecipeByID)
	router.PUT("/:id", UpdateRecipe)
	router.DELETE("/:id", DeleteRecipe)
}