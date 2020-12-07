package users


import (
	// "fmt"
	"os"
	"time"
	
	"go_server/Config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// Sets the user and the user id in session
func SetSessionUser(c *gin.Context, user_id uint){
	var userModel UserModel
	if user_id != 0 {
		Config.DB.First(&userModel, user_id)
	}

	c.Set("my_user_id", user_id)
	c.Set("my_user_model", userModel)

	// fmt.Println(c.MustGet("my_user_model"))
	// fmt.Println(c.MustGet("my_user_id"))
}


// Creates the user bearer
func CreateBearer(email string) string {
	var err error

	// Set claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate bearer with claims
	bearer := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded bearer using the secret signing key
	b, err := bearer.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "undefined"
	}
	return b
}