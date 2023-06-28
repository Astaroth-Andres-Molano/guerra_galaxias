// models/models.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type SatelliteData struct {
	gorm.Model
	ID            uint     `gorm:"primary_key"`
	SatelliteName string   `gorm:"not null"`
	Distance      float64  `gorm:"not null"`
	Message       []string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
