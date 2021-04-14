package main

import (
	"coinchain/helpers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


func main(){
	isValid := helpers.ValidateInteger(12)
	fmt.Println("Hello api", isValid)
	fmt.Printf("Redis host is %s and port %s\n", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
}