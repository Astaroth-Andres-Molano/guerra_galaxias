package helpers

import (
	"encoding/json"
	"guerra_galaxias/models"
)

func AppendSatellites(sat models.Satellite) []models.Satellite {
	var satel []models.Satellite

	satel = append(satel, sat)
	return satel
}

func FormatoMessage(message string) []string {
	var dataMessage []string
	var strMessage []byte
	strMessage = []byte(message)
	json.Unmarshal(strMessage, &dataMessage)
	return dataMessage
}
