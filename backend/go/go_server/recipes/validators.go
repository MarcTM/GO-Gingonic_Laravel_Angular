package recipes


import(
	// "fmt"

	"github.com/gin-gonic/gin"

	// "go_server/users"
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
	// myUserModel := c.MustGet("my_user_model").(users.UserModel)
	// fmt.Println(myUserModel)

	err := c.ShouldBindJSON(self)
	if err != nil {
		return err
	}

	return nil

	// s.articleModel.Slug = slug.Make(s.Article.Title)
	// s.articleModel.Title = s.Article.Title
	// s.articleModel.Description = s.Article.Description
	// s.articleModel.Body = s.Article.Body
	// s.articleModel.Author = GetArticleUserModel(myUserModel)
	// s.articleModel.setTags(s.Article.Tags)
	// return nil
}