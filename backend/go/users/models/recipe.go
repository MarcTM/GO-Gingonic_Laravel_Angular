package models

type RecipeModel struct {
	Id      	uint   `json:"id"`
	Name    	string `json:"name"`
	Description string `json:"description"`
	UserModelID	uint
	UserModel   models.UserModel
}

func (recipe *RecipeModel) TableName() string {
    return "recipes"
}