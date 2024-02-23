package models

import "time"

type User struct {
	ID        int       `json:"ID" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName" gorm:"unique"`
	LastName  *string   `json:"lastName"`
}
