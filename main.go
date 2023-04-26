package main

import (
	"log"

	"checkout-backend/app"
	"checkout-backend/app/database"
	"checkout-backend/app/routes"
	_ "checkout-backend/docs"
)

var instance *app.Server

func main() {
	mysql := database.InitMYSQL()
	instance = app.NewServer(mysql)
	routes.Routes(instance)
	err := instance.Start("8888")
	if err != nil {
		log.Fatal("Port is already occupied.")
	}
}
