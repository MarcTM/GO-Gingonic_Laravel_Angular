package users

import(
	"go_server/Config"
	"go_server/models"
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
	myUserModel := self.c.MustGet("my_user_model").(models.UserModel)
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


// Short profile serializer
type ShortProfileSerializer struct {
	profile models.UserModel
}

type ShortProfileResponse struct {
	ID		 uint	 `json:"id"`
	Username string  `json:"username"`
	Bio      string  `json:"bio"`
	Image    string  `json:"image"`
}

func (self *ShortProfileSerializer) Response() ShortProfileResponse{
	profile := ShortProfileResponse{
		ID:       self.profile.ID,
		Username: self.profile.Username,
		Bio:      self.profile.Bio,
		Image:    self.profile.Image,
	}
	return profile
}


// Profile with recipes serializer
type ProfileSerializer struct {
	profile models.UserModel
}

type ProfileResponse struct {
	ID		 uint	 `json:"id"`
	Username string  `json:"username"`
	Bio      string  `json:"bio"`
	Image    string  `json:"image"`
	Recipes  []ProfileRecipesResponse  `json:"recipes"`
}

func (self *ProfileSerializer) Response() ProfileResponse{
	var recipes []models.RecipeModel
	Config.DB.Model(self.profile).Association("Recipes").Find(&recipes)
	
	profileRecipesResponse := ProfileRecipesSerializer{recipes}
	profile := ProfileResponse{
		ID:       self.profile.ID,
		Username: self.profile.Username,
		Bio:      self.profile.Bio,
		Image:    self.profile.Image,
		Recipes:  profileRecipesResponse.Response(),
	}
	return profile
}


// Profile recipes serializer
type ProfileRecipesSerializer struct {
	recipes []models.RecipeModel
}

type ProfileRecipesResponse struct {
	ID		    uint	`json:"id"`
	Name     	string  `json:"name"`
	Description string  `json:"description"`
}

func (self *ProfileRecipesSerializer) Response() []ProfileRecipesResponse {
	var allrecipes []ProfileRecipesResponse
	for _, r := range self.recipes {
		onerecipe := ProfileRecipesResponse{
			ID:			 r.Id,
			Name:		 r.Name,
			Description: r.Description,
		}
		allrecipes = append(allrecipes, onerecipe)
	}
	return allrecipes
}