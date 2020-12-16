package users

import (
	"time"
	"go_server/Config"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)


// User model
type UserModel struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username" gorm:"not null;unique;type:varchar(191)"`
	Email        string  `json:"email" gorm:"not null;unique;type:varchar(191)"`
	Password     string  `json:"password" gorm:"not null;type:varchar(191)"`
	Bio          string  `json:"bio" gorm:"type:varchar(191)"`
	Image        string  `json:"image" gorm:"type:varchar(191)"`
	Type		 string  `json:"type" gorm:"not null;type:varchar(191)"`
	CreatedAt	 time.Time
	UpdatedAt	 time.Time
}

func (user *UserModel) TableName() string {
    return "users"
}


// Find user in database and return it
func FindUser(condition interface{}) (UserModel, error) {
	var model UserModel

	err := Config.DB.Where(condition).First(&model).Error
	return model, err
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


// Create hashed pass
func HashPass(str string) (string) {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
    return string(hashed)
}


// Compare password with hashed password
func UnhashPass(str string, hashed string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}