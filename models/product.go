package models

import "time"

type Product struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number" gorm:"unique"`
}
