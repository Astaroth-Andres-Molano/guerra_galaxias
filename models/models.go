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
	ID            uint     `gorm:"primary_key"`
	SatelliteName string   `gorm:"not null"`
	Distance      float64  `gorm:"not null"`
	Message       []string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Antes de guardar el registro, se serializa el slice de strings a JSON
// func (s *SatelliteData) BeforeSave(tx *gorm.DB) error {
// 	jsonData, err := json.Marshal(s.Message)
// 	if err != nil {
// 		return err
// 	}
// 	s.Message = jsonData
// 	return nil
// }

// Después de cargar el registro, se deserializa el JSON a slice de strings
// func (s *SatelliteData) AfterFind(tx *gorm.DB) error {
// 	var message []string
// 	if err := json.Unmarshal(s.Message, &message); err != nil {
// 		return err
// 	}
// 	s.Message = message
// 	return nil
// }
