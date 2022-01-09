package main

import (
	"gameSites/app/controllers"
	"log"
)

func main() {
	log.Fatalln(controllers.StartMainServer())
}
