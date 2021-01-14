package recipes

import(
	"go_server/models"
	"github.com/gin-gonic/gin"
)


// Recipe validator
func NewRecipeModelValidator() RecipeModelValidator {
	return RecipeModelValidator{}
}

type RecipeModelValidator struct {
	Recipe struct {
		Name        string   `form:"name" json:"name" binding:"required"`
		Description string   `form:"description" json:"description" binding:"required"`
		Image	    string	 `form:"image" json:"image" binding:"required"`
	} `json:"recipe"`
	recipeModel models.RecipeModel `json:"-"`
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


// Comment validator
func NewCommentModelValidator(id string) CommentModelValidator {
	return CommentModelValidator{Recipe: id}
}

type CommentModelValidator struct {
	Recipe  string
	Comment struct {
		Body		string	`form:"body" json:"body" binding:"required"`
	} `json:"comment"`
	commentModel models.CommentModel `json:"-"`
}

func (self *CommentModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(models.UserModel)

	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	var recipe models.RecipeModel
	Get(&recipe, self.Recipe)

	self.commentModel.Body = self.Comment.Body
	self.commentModel.UserModel = myUserModel
	self.commentModel.RecipeModel = recipe
	return nil
}