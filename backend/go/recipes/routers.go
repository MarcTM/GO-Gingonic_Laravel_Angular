package recipes

import (
	"go_server/users"

	"github.com/gin-gonic/gin"
)

// /recipes
func RecipesRegister(router *gin.RouterGroup) {
	router.GET("/", GetRecipes)
	router.POST("/", users.IsAuthenticated(CreateRecipe))
	router.GET("/:id", GetRecipeByID)
	router.PUT("/:id", UpdateRecipe)
	router.DELETE("/:id", DeleteRecipe)
}