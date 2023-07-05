// models/models.go

package models

import (
	//"encoding/json"
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
	ID            uint    `gorm:"primary_key"`
	SatelliteName string  `gorm:"type:character varying; uniqueIndex:idx_name; not null"`
	Distance      float64 `gorm:"type: numeric; not null"`
	Message       string  `gorm:"type: character varying; not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UpdateSatellite struct {
	SatelliteName string  `json:"name"`
	Distance      float64 `json:"distance"`
	Message       string  `json:"message"`
}
