package recipes

import (
	"github.com/gin-gonic/gin"
)


func RecipesRegister(router *gin.RouterGroup) {
	router.GET("/", GetRecipes)
	router.POST("/", CreateRecipe)
	router.GET("/:id", GetRecipeByID)
	router.PUT("/:id", UpdateRecipe)
	router.DELETE("/:id", DeleteRecipe)
}