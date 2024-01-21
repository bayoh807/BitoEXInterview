package main

import (
	_ "github.com/joho/godotenv/autoload"
	"tinder-server/routes"
	_ "tinder-server/routes"
)

func main() {

	routes.Routers.Run()

}
