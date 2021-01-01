package recipes

import (
	"fmt"
	"go_server/Config"
	_ "github.com/go-sql-driver/mysql"
	"go_server/models"
)


//GetAllRecipes
func GetAll(recipe *[]models.RecipeModel) (err error) {
	if err = Config.DB.Find(recipe).Error; err != nil {
		return err
	}
	return nil
}

//CreateRecipe ... Insert New data
func Create(recipe *models.RecipeModel) (err error) {
	if err = Config.DB.Create(recipe).Error; err != nil {
		return err
	}
	return nil
}

//GetRecipeByID ... Fetch only one recipe by Id
func Get(recipe *models.RecipeModel, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(recipe).Error; err != nil {
		return err
	}
	return nil
}

//UpdateRecipe ... Update recipe
func Update(recipe *models.RecipeModel, id string) (err error) {
	var update models.RecipeModel
	Config.DB.Where("id = ?", id).First(&update)

	Config.DB.Model(&update).Updates(recipe)
	return nil
}

//DeleteRecipe ... Delete recipe
func Delete(recipe *models.RecipeModel, id string) (err error) {
	if err = Config.DB.Where("ID = ?", id).First(recipe).Error; err != nil {
		return err
	} else {
		fmt.Println(recipe)
		fmt.Println(&recipe)
		fmt.Println(id)
		Config.DB.Where("id = ?", id).Delete(recipe)
		return nil
	}
}