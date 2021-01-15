package recipes

import(
	"net/http"
	"fmt"
	"go_server/Config"
	"go_server/models"
	"github.com/gin-gonic/gin"
)


// Short recipe serializer, without author
type ShortRecipeSerializer struct {
	recipe models.RecipeModel
}

type ShortRecipeResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Image		   string				 `json:"image"`
}

func (self *ShortRecipeSerializer) Response() ShortRecipeResponse {
	response := ShortRecipeResponse{
		ID:          self.recipe.Id,
		Name:        self.recipe.Name,
		Description: self.recipe.Description,
		Image:		 self.recipe.Image,
	}
	return response
}


// One recipe serializer
type OneRecipeSerializer struct {
	c *gin.Context
}

type OneRecipeResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Image		   string				 `json:"image"`
	Author         RecipeProfileResponse `json:"author"`
}

func (self *OneRecipeSerializer) Response() OneRecipeResponse {
	id := self.c.Params.ByName("id")
	var recipe models.RecipeModel
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
		Image:		 recipe.Image,
		Author:      recipeProfileSerializer.Response(),
	}
	return response
}


// All recipes serializer
type AllRecipeSerializer struct {
	c *gin.Context
}

type AllRecipeResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Image		   string				 `json:"image"`
	Author         RecipeProfileResponse `json:"author"`
}

func (self *AllRecipeSerializer) Response() []AllRecipeResponse {
	var recipe []models.RecipeModel
	err := GetAll(&recipe)
	if err != nil {
		fmt.Println("error")
		self.c.AbortWithStatus(http.StatusNotFound)
	}

	var recipeAll []AllRecipeResponse
	for _, r := range recipe {
		recipeProfileSerializer := RecipeProfileSerializer{r}

		response := AllRecipeResponse{
			ID:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Image:		 r.Image,
			Author:      recipeProfileSerializer.Response(),
		}
		recipeAll = append(recipeAll, response)
	}
	return recipeAll
}


// Recipe author profile serializer
type RecipeProfileSerializer struct {
	recipe models.RecipeModel
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

// Short comments serializer
type ShortCommentsSerializer struct {
	comments []models.CommentModel
}

type ShortCommentsResponse struct {
	ID             uint                  `json:"id"`
	Body           string                `json:"body"`
	RecipeModelID  uint                  `json:"recipe_id"`
	UserModelID    uint					 `json:"user_id"`
}

func (self *ShortCommentsSerializer) Response() []ShortCommentsResponse {
	var allcomments []ShortCommentsResponse
	for _, r := range self.comments {
		onecomment := ShortCommentsResponse{
			ID:			   r.Id,
			Body:		   r.Body,
			RecipeModelID: r.RecipeModelID,
			UserModelID:   r.UserModelID,
		}
		allcomments = append(allcomments, onecomment)
	}
	return allcomments
}