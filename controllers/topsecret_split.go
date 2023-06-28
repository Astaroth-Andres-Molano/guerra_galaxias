// controllers/topsecret_split.go

package controllers

import (
	"net/http"

	"guerra_galaxias/db"
	"guerra_galaxias/models"

	"github.com/gin-gonic/gin"
)

// type Satellite = models.Satellite
// type Position = models.Position
// type Response = models.Response
type SatelliteData = models.SatelliteData

func SaveSatelliteData(c *gin.Context) {
	satelliteName := c.Param("satellite_name")

	var payload struct {
		Distance float64  `json:"distance"`
		Message  []string `json:"message"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	satelliteData := SatelliteData{
		SatelliteName: satelliteName,
		Distance:      payload.Distance,
		Message:       payload.Message,
	}

	err := db.DB.Create(&satelliteData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func GetTopSecretSplit(c *gin.Context) {
	var satellitesData []SatelliteData

	err := db.DB.Find(&satellitesData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(satellitesData) != 3 {
		c.JSON(http.StatusNotFound, gin.H{"error": "insufficient satellite data"})
		return
	}

	var satellites []Satellite
	for _, data := range satellitesData {
		satellite := Satellite{
			Name:     data.SatelliteName,
			Distance: data.Distance,
			Message:  data.Message,
		}
		satellites = append(satellites, satellite)
	}

	position, err := GetLocation(satellites)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	message := GetMessage(satellites)

	response := Response{
		Position: position,
		Message:  message,
	}

	c.JSON(http.StatusOK, response)
}
