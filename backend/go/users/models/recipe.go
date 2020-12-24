package models

type RecipeModel struct {
	Id      	uint   `json:"id"`
	Name    	string `json:"name"`
	Description string `json:"description"`
	AuthorID	uint   `json:"author_id"`
}

func (recipe *RecipeModel) TableName() string {
    return "recipes"
}