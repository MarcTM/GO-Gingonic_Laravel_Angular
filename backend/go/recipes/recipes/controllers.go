package recipes

import (
	"net/http"
	"go_server/models"
	"github.com/gin-gonic/gin"
)


//CreateRecipe
func CreateRecipe(c *gin.Context) {
	recipeModelValidator := NewRecipeModelValidator()
	if err := recipeModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
		return
	}

	if err := Create(&recipeModelValidator.recipeModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "error")
		return
	}

	c.JSON(http.StatusCreated, "ok")
}


//GetRecipes
func GetRecipes(c *gin.Context) {
	serializer := AllRecipeSerializer{c}
	recipes := serializer.Response()
	c.JSON(http.StatusOK, recipes)
}


//GetRecipeByID
func GetRecipeByID(c *gin.Context) {
	serializer := OneRecipeSerializer{c}
	recipe := serializer.Response()
	c.JSON(http.StatusOK, recipe)
}


//UpdateRecipe
func UpdateRecipe(c *gin.Context) {
	var recipe models.RecipeModel
	id := c.Params.ByName("id")
	err := Get(&recipe, id)
	if err != nil {
		c.JSON(http.StatusNotFound, recipe)
	} else{
		err = Update(&recipe, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, recipe)
		}
	}
}


//DeleteRecipe
func DeleteRecipe(c *gin.Context) {
	var recipe models.RecipeModel
	id := c.Params.ByName("id")
	err := Delete(&recipe, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}