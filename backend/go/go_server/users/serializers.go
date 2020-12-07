package users


import(
	"github.com/gin-gonic/gin"
)


// Users serializers
type UserSerializer struct {
	c *gin.Context
}


type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    string  `json:"image"`
	Bearer   string  `json:"bearer"`
}


func (self *UserSerializer) Response() UserResponse{
	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Bearer:    CreateBearer(myUserModel.Email),
	}
	return user
}