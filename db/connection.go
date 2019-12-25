package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDB(dbDialect string, dbHost string, dbPort string, dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbUser, dbPort, dbPassword, dbName,
	)
	fmt.Println(dbDialect, connectionString)

	return gorm.Open(dbDialect, connectionString)
}
