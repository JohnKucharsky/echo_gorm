package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ProductID int       `json:"product_id" gorm:"uniqueIndex:product_user"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	UserID    int       `json:"user_id" gorm:"uniqueIndex:product_user"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
