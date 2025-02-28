package main

import (
	"context"
	"go_learning/db"
)

func main() {
	Run()
}

func Run() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	os.Exit(1)
	// }

	db.DBConn = SetupPGX()
	defer db.DBConn.Close(context.Background())

	router := SetupRouter()

	router.Run(":8080")
}
