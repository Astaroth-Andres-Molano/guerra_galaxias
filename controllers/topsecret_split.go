// controllers/topsecret_split.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"guerra_galaxias/db"
	"guerra_galaxias/helpers"
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

	message, _ := json.Marshal(payload.Message)
	messageStr := string(message)
	fmt.Println(messageStr, "messageData")
	satelliteData := SatelliteData{
		SatelliteName: satelliteName,
		Distance:      payload.Distance,
		Message:       messageStr,
	}

	err := db.DB.Create(&satelliteData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
	c.JSON(http.StatusOK, gin.H{
		"response":   "Registro creado con éxito!",
		"StatusCode": http.StatusOK,
	})
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

	var dataMessage []string
	var satellites []Satellite
	for _, data := range satellitesData {
		dataMessage = helpers.FormatoMessage(data.Message)
		satellite := Satellite{
			Name:     data.SatelliteName,
			Distance: data.Distance,
			Message:  dataMessage,
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

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": http.StatusOK,
		"response":   response,
		//"satellites": satellites,
	})
}

func GetOneSatellite(c *gin.Context) {

	satellite := helpers.GetOneS(c.Param("satellite_name"))

	if satellite == nil {
		c.JSON(http.StatusNotFound, gin.H{"StatusCode": http.StatusNotFound, "error": "Datos no encontrados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response":   satellite,
		"StatusCode": http.StatusOK,
		//"satellites": satellites,
	})
}

func GetAllSatellites(c *gin.Context) {
	var satellitesData []SatelliteData

	err := db.DB.Find(&satellitesData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": http.StatusOK,
		"response":   satellitesData,
	})
}

func UpdateSatellite(c *gin.Context) {

	satellite := helpers.GetOneS(c.Param("satellite_name"))

	if satellite == nil {
		c.JSON(http.StatusNotFound, gin.H{"StatusCode": http.StatusNotFound, "error": "Datos no encontrados"})
		return
	}

	var payload models.Satellite

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageStr := helpers.FormatoSendMessage(payload.Message)
	satelliteData := SatelliteData{
		SatelliteName: payload.Name,
		Distance:      payload.Distance,
		Message:       messageStr,
	}
	db.DB.Model(&satellite).Updates(satelliteData)

	c.JSON(http.StatusOK, gin.H{"data": satellite, "message": "Datos actualizados correctamente!"})
}

func DeleteSatellite(c *gin.Context) {
	satellite := helpers.GetOneS(c.Param("satellite_name"))

	if satellite == nil {
		c.JSON(http.StatusNotFound, gin.H{"StatusCode": http.StatusNotFound, "error": "Datos no encontrados"})
		return
	}

	if err := db.DB.Delete(&satellite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"StatusCode": http.StatusOK, "message": "Satellite deleted"})
}
