package main

import (
	"final-project/database"
	"final-project/routers"
	"fmt"
)

func main() {
	database.StartDB()
	routers.StartApp().Run()

	fmt.Println("Successfully connected to database")
}
