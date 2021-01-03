package models

type RecipeModel struct {
	Id      	uint   `json:"id"`
	Name    	string `json:"name"`
	Description string `json:"description"`
	Image		string `json:"image"`
	UserModelID	uint
	UserModel	UserModel
}

func (recipe *RecipeModel) TableName() string {
    return "recipes"
}