package users


import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)


//Get users
func GetUsers(c *gin.Context) {
	var user []UserModel
	err := GetAll(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


// Register user
func RegisterUser(c *gin.Context) {
	registerValidator := NewRegisterValidator()

	if err := registerValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect JSON")
		return
	}

	// Checks if email and password exist in database, if not, creates the user
	if err := CheckUserRegister(&registerValidator.userModel, registerValidator.userModel.Username, registerValidator.userModel.Email); err != nil {
		if err := Create(&registerValidator.userModel); err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, registerValidator.userModel)
		}
		return
	}
	c.JSON(http.StatusUnprocessableEntity, "The username or email is already token")
}


// Login user
func LoginUser(c *gin.Context) {
	loginValidator := NewLoginValidator()

	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect JSON")
		return
	}

	userModel, err := FindUser(&UserModel{Email: loginValidator.userModel.Email})
	
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Email doesn't exist")
		return
	}

	if user := CheckUser(&userModel, loginValidator.userModel.Email, loginValidator.userModel.Password); user == false {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect password")
		return
	}

	SetSessionUser(c, userModel.ID)

	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}