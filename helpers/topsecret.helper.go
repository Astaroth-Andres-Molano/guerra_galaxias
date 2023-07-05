package helpers

import (
	"encoding/json"
	"guerra_galaxias/db"
	"guerra_galaxias/models"
)

func AppendSatellites(sat models.Satellite) []models.Satellite {
	var satel []models.Satellite

	satel = append(satel, sat)
	return satel
}

func FormatoMessage(message string) []string {
	var dataMessage []string
	strMessage := []byte(message)
	json.Unmarshal(strMessage, &dataMessage)
	return dataMessage
}

func FormatoSendMessage(message []string) string {
	dataMessage, _ := json.Marshal(message)
	messageStr := string(dataMessage)
	return messageStr
}

func GetOneS(nameSatellite string) (satellite []models.SatelliteData) {
	//var satellite []models.SatelliteData
	if err := db.DB.Where("satellite_name = ?", nameSatellite).First(&satellite).Error; err != nil {
		return nil
	}

	return satellite
}
