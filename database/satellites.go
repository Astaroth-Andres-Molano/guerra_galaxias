package database

type Location struct {
	X float32
	Y float32
}

var (
	kenobiLocation    = Location{-500, -200}
	skywalkerLocation = Location{100, -100}
	satoLocation      = Location{500, 100}

	Satellites = map[string]Location{
		"kenobi":    kenobiLocation,
		"skywalker": skywalkerLocation,
		"sato":      satoLocation,
	}

	SatelliteNames = []string{"kenobi", "skywalker", "sato"}
)
