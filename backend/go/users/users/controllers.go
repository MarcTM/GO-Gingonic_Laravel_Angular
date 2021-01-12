package users

import (
	"net/http"
	"fmt"
	"go_server/Config"
	"go_server/models"
	"github.com/gin-gonic/gin"
)

// Get your account profile
func GetAccount(c *gin.Context) {

}


// Get user profile
func GetProfile(c *gin.Context) {
	username := c.Params.ByName("username")
	userModel, err := FindUser(&models.UserModel{Username: username})

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "User not found")
		return
	}

	profileSerializer := ProfileSerializer{userModel}
	c.JSON(http.StatusOK, profileSerializer.Response())
}


//Get users
func GetUsers(c *gin.Context) {
	var user []models.UserModel

	err := GetAll(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


//Me
func Me(c *gin.Context) {
	me := c.MustGet("my_user_id")

	var myprofile models.UserModel
	Config.DB.Where("id = ?", me).First(&myprofile)

	meSerializer := ShortProfileSerializer{myprofile}
	c.JSON(http.StatusOK, meSerializer.Response())
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

	userModel, err := FindUser(&models.UserModel{Email: loginValidator.userModel.Email})
	
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


// // Update profile
// func UpdateProfile(c *gin.Context) {
// 	id := c.Params.ByName("id")

// 	recipeModelValidator := NewRecipeModelValidator()
// 	if err := recipeModelValidator.Bind(c); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
// 		return
// 	}

// 	fmt.Println(&recipeModelValidator.recipeModel)
// 	err := Update(&recipeModelValidator.recipeModel, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, "ok")
// 	}
// }


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


// Already following?
func IsFollowing(c *gin.Context) {
	username := c.Params.ByName("username")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Where("id = ?", myUserModel.ID).First(&user)

	var followers models.UserModel
	err := Config.DB.Model(user).Where("username = ?", username).Association("Following").Find(&followers).Error

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Not found")
	} else {
		c.JSON(http.StatusOK, "Found")
	}
}


// Follow user
func FollowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Preload("Following").First(&user, "id = ?", myUserModel.ID)
	
	var followed models.UserModel
	Config.DB.Where("id = ?", id).First(&followed)

	Config.DB.Model(user).Association("Following").Append(followed)
	c.JSON(http.StatusOK, "OK")
}


// Unfollow user
func UnfollowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	var user models.UserModel
	Config.DB.Preload("Following").First(&user, "id = ?", myUserModel.ID)
	
	var unfollowed models.UserModel
	Config.DB.Where("id = ?", id).First(&unfollowed)

	Config.DB.Model(user).Association("Following").Delete(unfollowed)
	c.JSON(http.StatusOK, "OK")
}


// // Count following
// func Count(c *gin.Context) {
// 	myUserModel := c.MustGet("my_user_model").(models.UserModel)

// 	var user models.UserModel
// 	Config.DB.Preload("Following").First(&user, "id = ?", myUserModel.ID)

// 	count := Config.DB.Model(user).Association("Following").Count()
// 	fmt.Println(count)
// }