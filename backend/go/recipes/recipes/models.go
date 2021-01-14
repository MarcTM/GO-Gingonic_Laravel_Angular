package recipes

import (
	"fmt"
	"go_server/Config"
	_ "github.com/go-sql-driver/mysql"
	"go_server/models"
)


//Get all recipes
func GetAll(recipe *[]models.RecipeModel) (err error) {
	if err = Config.DB.Order("created_at desc").Find(recipe).Error; err != nil {
		return err
	}
	return nil
}


//Create recipe
func Create(recipe *models.RecipeModel) (err error) {
	if err = Config.DB.Create(recipe).Error; err != nil {
		return err
	}
	return nil
}


//Get recipe by ID
func Get(recipe *models.RecipeModel, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(recipe).Error; err != nil {
		return err
	}
	return nil
}


//Update recipe
func Update(recipe *models.RecipeModel, id string) (err error) {
	var update models.RecipeModel
	Config.DB.Where("id = ?", id).First(&update)

	Config.DB.Model(&update).Updates(recipe)
	return nil
}


//Delete recipe
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


// Save comment
func SaveComment(comment *models.CommentModel) (err error) {
	if err = Config.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}