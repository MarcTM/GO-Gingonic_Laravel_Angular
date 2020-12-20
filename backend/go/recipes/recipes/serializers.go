package recipes


import(
	"go_server/models"
	"github.com/gin-gonic/gin"
)


// Recipe serializer
type RecipeSerializer struct {
	c *gin.Context
	RecipeModel
}


type RecipeResponse struct {
	ID             uint                  `json:"-"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Author         ProfileResponse		 `json:"author"`
}


func (self *RecipeSerializer) Response() RecipeResponse {
	profileSerializer := ProfileSerializer{self.c}

	response := RecipeResponse{
		ID:          self.Id,
		Name:        self.Name,
		Description: self.Description,
		Author:      profileSerializer.Response(),
	}
	return response
}



// Profile serializers
type ProfileSerializer struct {
	c *gin.Context
}


type ProfileResponse struct {
	ID		 uint	 `json:"id"`
	Username string  `json:"username"`
	Bio      string  `json:"bio"`
	Image    string  `json:"image"`
}


func (self *ProfileSerializer) Response() ProfileResponse{
	myUserModel := self.c.MustGet("my_user_model").(models.UserModel)
	profile := ProfileResponse{
		ID:		  myUserModel.ID,
		Username: myUserModel.Username,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
	}
	return profile
}