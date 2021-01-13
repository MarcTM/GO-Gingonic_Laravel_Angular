package users

import (
	"go_server/models"
	"github.com/gin-gonic/gin"
)


// Register validator
func NewRegisterValidator() RegisterValidator {
	registerValidator := RegisterValidator{}
	return registerValidator
}

type RegisterValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (self *RegisterValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	self.userModel.Username = self.User.Username
	self.userModel.Email = self.User.Email
	self.userModel.Password = HashPass(self.User.Password)
	self.userModel.Type = "user"

	return nil
}


// Login validator
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	self.userModel.Email = self.User.Email
	self.userModel.Password = self.User.Password
	return nil
}


// Update profile validator
func NewUpdateProfileValidator() UpdateProfileValidator {
	updateProfileValidator := UpdateProfileValidator{}
	return updateProfileValidator
}

type UpdateProfileValidator struct {
	User struct {
		Image    string `form:"image" json:"image"`
		Bio 	 string `form:"bio" json:"bio"`
	} `json:"user"`
	userModel models.UserModel `json:"-"`
}

func (self *UpdateProfileValidator) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	self.userModel.Image = self.User.Image
	self.userModel.Bio = self.User.Bio

	return nil
}