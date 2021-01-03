package recipes


import(
	"go_server/models"
	"github.com/gin-gonic/gin"
)


type RecipeModelValidator struct {
	Recipe struct {
		Name        string   `form:"name" json:"name" binding:"required"`
		Description string   `form:"description" json:"description" binding:"required"`
		Image	    string	 `form:"image" json:"image" binding:"required"`
	} `json:"recipe"`
	recipeModel models.RecipeModel `json:"-"`
}


func NewRecipeModelValidator() RecipeModelValidator {
	return RecipeModelValidator{}
}


func (self *RecipeModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	self.recipeModel.Name = self.Recipe.Name
	self.recipeModel.Description = self.Recipe.Description
	self.recipeModel.Image = self.Recipe.Image
	self.recipeModel.UserModel = myUserModel
	return nil
}