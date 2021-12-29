package main

import (
	"log"
	"todo_app/todo_app/app/controllers"
)

func main() {
	log.Fatalln(controllers.StartMainServer())
}
