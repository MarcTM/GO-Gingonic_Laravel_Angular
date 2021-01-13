package users

import (
	"go_server/models"
	"go_server/Config"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)


// Find user in database and return it
func FindUser(condition interface{}) (models.UserModel, error) {
	var model models.UserModel

	err := Config.DB.Where(condition).First(&model).Error
	return model, err
}


// Update user (profile)
func Update(user *models.UserModel, id uint) (err error) {
	var update models.UserModel
	Config.DB.Where("id = ?", id).First(&update)

	Config.DB.Model(&update).Updates(user)
	return nil
}


// Check if user exists by email or username (register)
func CheckUserRegister(user *models.UserModel, username string, email string) (error) {
	err := Config.DB.Where("email = ?", email).Or("username = ?", username).First(user).Error

	if err != nil{
		return err
	}
	return nil
}


//Get all users
func GetAll(user *[]models.UserModel) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}


//Register user
func Create(user *models.UserModel) (err error) {
	user.Image = "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/default-avatar.png"
	user.Bio = "(This user has no bio yet)"
	
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