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
	Type	 string  `json:"type"`
}

func (self *UserSerializer) Response() UserResponse{

	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Bearer:   CreateBearer(myUserModel.Email),
		Type:	  myUserModel.Type,
	}
	return user
}


// Profile serializers
type ProfileSerializer struct {
	profile UserModel
}

type ProfileResponse struct {
	ID		 uint	 `json:"id"`
	Username string  `json:"username"`
	Bio      string  `json:"bio"`
	Image    string  `json:"image"`
}

func (self *ProfileSerializer) Response() ProfileResponse{

	profile := ProfileResponse{
		ID:       self.profile.ID,
		Username: self.profile.Username,
		Bio:      self.profile.Bio,
		Image:    self.profile.Image,
	}
	return profile
}