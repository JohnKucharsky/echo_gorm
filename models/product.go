package models

import "time"

type Product struct {
	ID           int       `json:"ID" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"createdAt"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serialNumber" gorm:"unique"`
}
