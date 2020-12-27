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
	Config.DB.AutoMigrate(&models.UserModel{})

	// Set routes
	r := Routes.SetupRouter()

	// // Chuleta has many (association)
	// 
	// var user models.UserModel
	// Config.DB.Where("username = ?", "marc").First(&user)
	// fmt.Println(user)
	// var recipes []models.RecipeModel
	// Config.DB.Model(user).Association("Recipes").Find(&recipes)
	// fmt.Println(recipes)


	// Run application on port 3001
	r.Run(":3001")
}