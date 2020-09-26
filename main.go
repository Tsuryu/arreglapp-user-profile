package main

import (
	"fmt"

	"github.com/Tsuryu/arreglapp-user-profile/app/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: " + err.Error())
	}
	routers.Router()
}
