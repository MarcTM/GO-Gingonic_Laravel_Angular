package recipes


import(
	"net/http"
	"fmt"
	"go_server/Config"
	"go_server/models"
	"github.com/gin-gonic/gin"
)


// Recipe serializer when finding one
type OneRecipeSerializer struct {
	c *gin.Context
}

type OneRecipeResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Author         RecipeProfileResponse `json:"author"`
}

func (self *OneRecipeSerializer) Response() OneRecipeResponse {
	id := self.c.Params.ByName("id")
	var recipe RecipeModel
	err := Get(&recipe, id)	
	if err != nil {
		fmt.Println("error")
		self.c.AbortWithStatus(http.StatusNotFound)
	}
	recipeProfileSerializer := RecipeProfileSerializer{recipe}

	response := OneRecipeResponse{
		ID:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
		Author:      recipeProfileSerializer.Response(),
	}
	return response
}


// Profile for one recipe
type RecipeProfileSerializer struct {
	recipe RecipeModel
}

type RecipeProfileResponse struct {
	ID		 uint	 `json:"id"`
	Username string  `json:"username"`
	Image    string  `json:"image"`
}

func (self *RecipeProfileSerializer) Response() RecipeProfileResponse{
	var user models.UserModel
	Config.DB.Model(self.recipe).Related(&user)

	profile := RecipeProfileResponse{
		ID:		  user.ID,
		Username: user.Username,
		Image:    user.Image,
	}
	return profile
}