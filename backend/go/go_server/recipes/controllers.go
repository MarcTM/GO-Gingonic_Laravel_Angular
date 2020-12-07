package recipes

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)


//GetRecipes
func GetRecipes(c *gin.Context) {
	var recipe []RecipeModel
	err := GetAll(&recipe)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}

//CreateRecipe
func CreateRecipe(c *gin.Context) {
	var recipe RecipeModel
	c.BindJSON(&recipe)

	err := Create(&recipe)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}


//GetRecipeByID
func GetRecipeByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var recipe RecipeModel
	err := Get(&recipe, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, recipe)
	}
}

//UpdateRecipe
func UpdateRecipe(c *gin.Context) {
	var recipe RecipeModel
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
	var recipe RecipeModel
	id := c.Params.ByName("id")
	err := Delete(&recipe, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}