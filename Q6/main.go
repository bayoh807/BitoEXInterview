package main

import (
	_ "github.com/joho/godotenv/autoload"
	"tinder-Server/routes"
	_ "tinder-Server/routes"
)

func main() {

	routes.Routers.Run()

}
