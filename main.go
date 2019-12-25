package main

import (
	"log"
	"os"

	"github.com/burhanuddinbhopalwala/go-restful-backend-app/config"
	"github.com/burhanuddinbhopalwala/go-restful-backend-app/db"
	configs "github.com/burhanuddinbhopalwala/go-restful-backend-app/routes"
)

func main() {
	//* Loading the env variables
	config.LoadEnvVariables()

	//*  Database config
	db, err := db.ConnectToDB(os.Getenv("dbDialect"), os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))

	//*  Database connection error
	if err != nil {
		log.Fatalln(err)
	}

	//*  Ping to database
	err = db.DB().Ping()

	//*  Error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	//*  Migration
	//*  db.AutoMigrate(&models.Contact{})

	//* DB connection close
	defer db.Close()

	contactRepository := controllers.NewContactRepository(db)
	route := configs.SetupRoutes(contactRepository)
	route.Run(":8000")
}
