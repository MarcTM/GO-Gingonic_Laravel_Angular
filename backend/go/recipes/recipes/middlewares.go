package recipes


import (
	"strconv"
	"fmt"
	"os"
	"go_server/models"
	"go_server/Config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// Sets the user and the user id in session
func SetSessionUser(c *gin.Context, user_id uint){
	var userModel models.UserModel
	if user_id != 0 {
		Config.DB.First(&userModel, user_id)
	}
	fmt.Println(userModel.Email)

	c.Set("my_user_id", user_id)
	c.Set("my_user_model", userModel)
	// fmt.Println(c.MustGet("my_user_model"))
	// fmt.Println(c.MustGet("my_user_id"))
	Config.RedisSet("user_id", strconv.FormatUint(uint64(user_id), 10))
	Config.RedisSet("user_email", userModel.Email)
}


// Validate if there is a bearer in the headers, and if it is valid. Then, proceeds with the callback
func IsAuthenticated(endpoint func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Authorization"] != nil {

			const BEARER_SCHEMA = "Bearer "
			authHeader := c.GetHeader("Authorization")
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := ValidateToken(tokenString)
			if token.Valid {
				var user models.UserModel
				claims := token.Claims.(jwt.MapClaims)
				
				err := Config.DB.Where("email = ?", claims["email"]).First(&user).Error
				if err != nil {
					fmt.Println(err)
					return
				} else{
					SetSessionUser(c, user.ID)
					endpoint(c)
					return
				}
			} else {
				fmt.Println(err)
				fmt.Println("invalid")
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}
}


// Validates the received token and returns the claims
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
}