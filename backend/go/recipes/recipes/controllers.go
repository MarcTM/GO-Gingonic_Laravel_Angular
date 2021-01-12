package recipes

import (
	"fmt"
	"net/http"
	"go_server/models"
	"go_server/Config"
	"github.com/gin-gonic/gin"
)


//Create recipe
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


//Get recipes
func GetRecipes(c *gin.Context) {
	serializer := AllRecipeSerializer{c}
	recipes := serializer.Response()
	c.JSON(http.StatusOK, recipes)
}


//Get recipe by ID
func GetRecipeByID(c *gin.Context) {
	serializer := OneRecipeSerializer{c}
	recipe := serializer.Response()
	c.JSON(http.StatusOK, recipe)
}


//Update recipe
func UpdateRecipe(c *gin.Context) {
	id := c.Params.ByName("id")

	recipeModelValidator := NewRecipeModelValidator()
	if err := recipeModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
		return
	}

	fmt.Println(&recipeModelValidator.recipeModel)
	err := Update(&recipeModelValidator.recipeModel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "ok")
	}
}


//Delete recipe
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


// Already favorited?
func IsFavorited(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Where("id = ?", myUserModel.ID).First(&user)

	var favorites models.RecipeModel
	err := Config.DB.Model(user).Where("id = ?", id).Association("Favorites").Find(&favorites).Error

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Not found")
	} else {
		c.JSON(http.StatusOK, "Found")
	}
}


// Favorite recipe
func FavoriteRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Preload("Favorites").First(&user, "id = ?", myUserModel.ID)
	
	var favorited models.RecipeModel
	Config.DB.Where("id = ?", id).First(&favorited)

	Config.DB.Model(user).Association("Favorites").Append(favorited)
	c.JSON(http.StatusOK, "OK")
}


// Unfavorite recipe
func UnfavoriteRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Preload("Favorites").First(&user, "id = ?", myUserModel.ID)
	
	var unfavorited models.RecipeModel
	Config.DB.Where("id = ?", id).First(&unfavorited)

	Config.DB.Model(user).Association("Favorites").Delete(unfavorited)
	c.JSON(http.StatusOK, "OK")
}


// Return recipe if the user owns it
func OwnsRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Where("id = ?", myUserModel.ID).First(&user)

	var recipe models.RecipeModel
	err := Config.DB.Model(user).Where("id = ?", id).Association("Recipes").Find(&recipe).Error

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Not found")
	}

	response := ShortRecipeSerializer{recipe}
	c.JSON(http.StatusOK, response.Response())
}