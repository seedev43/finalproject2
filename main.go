package main

import (
	"fp2/database"
	"fp2/routers"
)

func main() {
	PORT := ":8080"

	_, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	routers.SetupRouter().Run(PORT)

}
