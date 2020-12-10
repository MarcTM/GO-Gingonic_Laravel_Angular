package main

import (
	"go_server/Config"
	"go_server/Routes"
	"go_server/users"
	"go_server/recipes"

	"fmt"

	"github.com/gin-contrib/cors"
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
	Config.DB.AutoMigrate(&users.UserModel{})
	Config.DB.AutoMigrate(&recipes.RecipeModel{})

	// Set routes
	r := Routes.SetupRouter()

	// CORS
	r.Use(cors.Default())
	// cors := func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	// 	c.Writer.Header().Set("Content-Type", "application/json")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(200)
	// 	}
	// 	c.Next()
	// }
	// r.Use(cors)

	// Run application on port 3000
	r.Run(":3000")
}
