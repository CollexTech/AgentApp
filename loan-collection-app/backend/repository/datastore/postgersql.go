package datastore

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgeSQLConn *gorm.DB

func ConnectPostgeSQL() error {
	if PostgeSQLConn != nil {
		return nil
	}
	host := viper.GetString("POSTGRESQL_HOST")
	port := viper.GetString("POSTGRESQL_PORT")
	dbName := viper.GetString("POSTGRESQL_DATABASE_NAME")
	user := viper.GetString("POSTGRESQL_USERNAME")
	password := viper.GetString("POSTGRESQL_PASSWORD")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
		return errors.Join(err, errors.New("unable to connect to POSTGRESQLl db"))
	}
	PostgeSQLConn = db
	return nil
}
