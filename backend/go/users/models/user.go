package models

import "time"

type UserModel struct {
	ID           uint    		`json:"id"`
	Username     string  		`json:"username" gorm:"not null;unique;type:varchar(191)"`
	Email        string  		`json:"email" gorm:"not null;unique;type:varchar(191)"`
	Password     string  		`json:"password" gorm:"not null;type:varchar(191)"`
	Bio          string  		`json:"bio" gorm:"type:varchar(191)"`
	Image        string  		`json:"image" gorm:"type:varchar(191)"`
	Type		 string  		`json:"type" gorm:"not null;type:varchar(191)"`
	Recipes		 []RecipeModel
	Following	 []UserModel	`gorm:"many2many:followers;association_jointable_foreignkey:following_id"`
	CreatedAt	 time.Time
	UpdatedAt	 time.Time
}

func (user *UserModel) TableName() string {
    return "users"
}