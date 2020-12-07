package recipes


import(
	"github.com/gin-gonic/gin"

	"go_server/users"
)


type RecipeModelValidator struct {
	Recipe struct {
		Name       string   `form:"name" json:"name" binding:"required"`
		Description string   `form:"description" json:"description" binding:"required"`
	} `json:"recipe"`
	recipeModel RecipeModel `json:"-"`
}


func NewRecipeModelValidator() RecipeModelValidator {
	return RecipeModelValidator{}
}


func (self *RecipeModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	self.recipeModel.Name = self.Recipe.Name
	self.recipeModel.Description = self.Recipe.Description
	self.recipeModel.Author = myUserModel
	self.recipeModel.AuthorID = myUserModel.ID
	return nil
}