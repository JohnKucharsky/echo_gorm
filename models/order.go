package models

import "time"

type Order struct {
	ID        int       `json:"ID" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	ProductID int       `json:"productID"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	UserID    int       `json:"userID"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
