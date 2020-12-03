package recipes

import (
	"fmt"
	
	"go_server/Config"

	_ "github.com/go-sql-driver/mysql"
)


type RecipeModel struct {
	Id      	uint   `gorm:"primary_key"`
	Name    	string `gorm;"column:name"`
	Description string `gorm:"column:description"`
	// Author		User   `gorm:"foreignKey:UserId" json:"author" binding:"required"`
	// UserId		int	   `json:"-"`
}


//GetAllRecipes
func GetAll(recipe *[]RecipeModel) (err error) {
	if err = Config.DB.Find(recipe).Error; err != nil {
		return err
	}
	return nil
}

//CreateRecipe ... Insert New data
func Create(recipe *RecipeModel) (err error) {
	if err = Config.DB.Create(recipe).Error; err != nil {
		return err
	}
	return nil
}

//GetRecipeByID ... Fetch only one recipe by Id
func Get(recipe *RecipeModel, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(recipe).Error; err != nil {
		return err
	}
	return nil
}

//UpdateRecipe ... Update recipe
func Update(recipe *RecipeModel, id string) (err error) {
	fmt.Println(recipe)
	Config.DB.Save(recipe)
	return nil
}

//DeleteRecipe ... Delete recipe
func Delete(recipe *RecipeModel, id string) (err error) {
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