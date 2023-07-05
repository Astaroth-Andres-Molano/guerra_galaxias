// main.go

package main

import (
	"fmt"
	"guerra_galaxias/controllers"
	"guerra_galaxias/db"
	"guerra_galaxias/models"

	"github.com/gin-gonic/gin"
)

type Satellite = models.Satellite
type Position = models.Position
type Response = models.Response

func main() {
	err := db.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	err = db.MigrateDB()
	if err != nil {
		fmt.Println("Failed to migrate the database:", err)
		return
	}

	router := gin.Default()

	router.POST("/topsecret/", controllers.GetTopSecret)
	router.POST("/topsecret_split/:satellite_name", controllers.SaveSatelliteData)
	router.GET("/topsecret_split/", controllers.GetTopSecretSplit)
	router.GET("/topsecret_split/:satellite_name", controllers.GetOneSatellite)
	router.GET("/topsecret_datasatellites/", controllers.GetAllSatellites)
	router.PUT("/topsecret_split/:satellite_name", controllers.UpdateSatellite)
	router.DELETE("/topsecret_split/:satellite_name", controllers.DeleteSatellite)

	router.Run()
}
