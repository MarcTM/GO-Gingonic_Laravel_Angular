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

			c.JSON(http.StatusOK, "Registration successfull")
		}
		return
	}
	c.JSON(http.StatusUnprocessableEntity, "The username or email is already taken")
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

	userpass := UnhashPass(loginValidator.User.Password, userModel.Password)
	if userpass == false {
		c.JSON(http.StatusUnprocessableEntity, "Incorrect password")
		return
	}

	SetSessionUser(c, userModel.ID)

	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}


// Validate if a token is correct or incorrect
func ValidateUserToken(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]

	token, err := ValidateToken(tokenString)

	if token.Valid {
		c.JSON(http.StatusOK, "Valid Token")
	} else {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, "Invalid Token")
	}
}