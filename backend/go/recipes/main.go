package main

import (
	"fmt"
	"go_server/models"
	"go_server/Config"
	"go_server/Routes"
	"github.com/jinzhu/gorm"
)


var err error

func main() {
	// Start connection to DB(mySql)
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	// Migrate tables
	Config.DB.AutoMigrate(&models.RecipeModel{})

	// Set routes
	r := Routes.SetupRouter()


	// // Chuleta has one (related)
	// 
	// var fRecipe models.RecipeModel
	// var use models.UserModel
	// Config.DB.Last(&fRecipe, "Name = ?", "p")
	// fmt.Println(fRecipe)
	// Config.DB.Model(fRecipe).Related(&use)
	// fmt.Println(use)


	// Run application on port 3000
	r.Run(":3002")
}