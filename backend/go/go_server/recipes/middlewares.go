package recipes


import(
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go_server/users"
	"go_server/Config"
)


// Validate if there is a bearer in the headers, and if it is valid. Then, proceeds with the callback
func isAuthenticated(endpoint func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Authorization"] != nil {

			const BEARER_SCHEMA = "Bearer "
			authHeader := c.GetHeader("Authorization")
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := ValidateToken(tokenString)
			if token.Valid {
				var user users.UserModel
				claims := token.Claims.(jwt.MapClaims)
				
				err := Config.DB.Where("email = ?", claims["email"]).First(&user).Error
				if err != nil {
					fmt.Println(err)
					return
				} else{
					users.SetSessionUser(c, user.ID)
					endpoint(c)
					return
				}
			} else {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}
}


// Validates the received token and returns the user email
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
}