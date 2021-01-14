package models

import "time"


type CommentModel struct {
	Id      	  uint   `json:"id"`
	Body		  string `json:"body"`
	RecipeModelID uint
	UserModelID   uint
	UserModel     UserModel
	RecipeModel	  RecipeModel
	CreatedAt	  time.Time
	UpdatedAt	  time.Time
}

func (comment *CommentModel) TableName() string {
    return "comments"
}