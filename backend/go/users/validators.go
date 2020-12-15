package users


import (
	"github.com/gin-gonic/gin"
)


// Register validator
type RegisterValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required"`
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel UserModel `json:"-"`
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


func NewRegisterValidator() RegisterValidator {
	registerValidator := RegisterValidator{}
	return registerValidator
}



// Login validator
type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel UserModel `json:"-"`
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


func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}