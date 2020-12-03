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
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&users.UserModel{})
	Config.DB.AutoMigrate(&recipes.RecipeModel{})

	r := Routes.SetupRouter()
	r.Use(cors.Default())

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	
	r.Run(":3000")
}
