package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // pgsql lib
)

// var ormObject orm.Ormer
var db *gorm.DB

const (
	dbhost = "DB_HOST"
	dbport = "DB_PORT"
	dbuser = "DB_USER"
	dbpass = "DB_PASS"
	dbname = "DB_NAME"
)

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() *gorm.DB {
	config := loadConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

	return db

	// orm.RegisterDriver("postgres", orm.DRPostgres)
	// orm.RegisterDataBase("default", "postgres", psqlInfo)

	// orm.RegisterModel(new(model.User))
	// orm.RegisterModel(new(model.Employee))
	// orm.RegisterModel(new(model.EmployeeStatus))

	// ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
// func GetOrmObject() orm.Ormer {
// 	return ormObject
// }

func loadConfig() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name

	return conf
}
