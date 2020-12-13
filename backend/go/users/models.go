package users

import (
	// "fmt"
	"go_server/Config"

	_ "github.com/go-sql-driver/mysql"
)


// User model
type UserModel struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username" gorm:"not null;unique"`
	Email        string  `json:"email" gorm:"not null;unique"`
	Password     string  `json:"password"`
	Bio          string  `json:"bio"`
	Image        string  `json:"image"`
	Type		 string  `json:"type"`
}


// Find user in database and return it
func FindUser(condition interface{}) (UserModel, error) {
	var model UserModel

	err := Config.DB.Where(condition).First(&model).Error
	return model, err
}


// Check if user email and password match (login)
func CheckUser(user *UserModel, email string, password string) bool {
	err := Config.DB.Where("email = ? AND password = ?", email, password).First(user).Error
	if err != nil {
		return false
	} else {
		return true
	}
}


// Check if user exists by email or username (register)
func CheckUserRegister(user *UserModel, username string, email string) (error) {
	err := Config.DB.Where("email = ?", email).Or("username = ?", username).First(user).Error

	if err != nil{
		return err
	}
	return nil
}


//Get all users
func GetAll(user *[]UserModel) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}


//Register user
func Create(user *UserModel) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}