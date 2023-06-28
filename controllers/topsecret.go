// controllers/topsecret.go

package controllers

import (
	"fmt"
	"net/http"

	"guerra_galaxias/models"

	"math"

	"github.com/gin-gonic/gin"
)

type Satellite = models.Satellite
type Position = models.Position
type Response = models.Response

var kenobiPosition = Position{X: -500, Y: -200}
var skywalkerPosition = Position{X: 100, Y: -100}
var satoPosition = Position{X: 500, Y: 100}

func GetTopSecret(c *gin.Context) {
	var payload struct {
		Satellites []Satellite `json:"satellites"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position, err := GetLocation(payload.Satellites)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	message := GetMessage(payload.Satellites)

	response := Response{
		Position: position,
		Message:  message,
	}

	c.JSON(http.StatusOK, response)
}

func GetLocation(distances []Satellite) (Position, error) {
	if len(distances) != 3 {
		return Position{}, fmt.Errorf("insufficient satellite data")
	}

	x, y, err := calculateCoordinates(distances[0].Distance, distances[1].Distance, distances[2].Distance)
	if err != nil {
		return Position{}, err
	}

	return Position{X: x, Y: y}, nil
}

func calculateCoordinates(d1, d2, d3 float64) (float64, float64, error) {
	x1, y1 := kenobiPosition.X, kenobiPosition.Y
	x2, y2 := skywalkerPosition.X, skywalkerPosition.Y
	x3, y3 := satoPosition.X, satoPosition.Y

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := math.Pow(float64(d1), 2) - math.Pow(float64(d2), 2) - math.Pow(float64(x1), 2) + math.Pow(float64(x2), 2) - math.Pow(float64(y1), 2) + math.Pow(float64(y2), 2)
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := math.Pow(float64(d2), 2) - math.Pow(float64(d3), 2) - math.Pow(float64(x2), 2) + math.Pow(float64(x3), 2) - math.Pow(float64(y2), 2) + math.Pow(float64(y3), 2)

	x := float64((C*E - F*B) / (E*A - B*D))
	y := float64((C*D - A*F) / (B*D - A*E))

	return x, y, nil
}

func GetMessage(messages []Satellite) string {
	if len(messages) != 3 {
		return ""
	}

	messageLength := getMessageLength(messages)
	result := make([]string, messageLength)

	for _, satellite := range messages {
		for i, word := range satellite.Message {
			if result[i] == "" && word != "" {
				result[i] = word
			}
		}
	}

	return joinWords(result)
}

func getMessageLength(messages []Satellite) int {
	maxLength := 0

	for _, satellite := range messages {
		if len(satellite.Message) > maxLength {
			maxLength = len(satellite.Message)
		}
	}

	return maxLength
}

func joinWords(words []string) string {
	result := ""

	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}

	return result
}
