package users


import (
	"fmt"
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
	// b, err := bearer.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	b, err := bearer.SignedString([]byte("marc_secret"))
	if err != nil {
		return "undefined"
	}
	return b
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
				var user UserModel
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