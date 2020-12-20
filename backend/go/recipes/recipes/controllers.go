package recipes


import (
	"net/http"
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
	recipeModelValidator := NewRecipeModelValidator()
	if err := recipeModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
		return
	}

	if err := Create(&recipeModelValidator.recipeModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "error")
		return
	}

	serializer := RecipeSerializer{c, recipeModelValidator.recipeModel}
	c.JSON(http.StatusCreated, gin.H{"recipe": serializer.Response()})
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