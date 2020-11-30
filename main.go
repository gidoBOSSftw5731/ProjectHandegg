package main

import (
	"fmt"
	"os"

	"github.com/gidoBOSSftw5731/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// userCreds is a struct that contains the structure for the database that is imported into gorm
// If the field is not capitalized, it will not end up in the db
type userCreds struct {
	UID     string `gorm:"primaryKey"`
	Balance float64
}

var (
	// Config
	config = make(map[string]string)
	// global db object
	db *gorm.DB
)

func main() {
	log.SetCallDepth(4)
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// get required config vars from env
	for _, i := range []string{"DBUSER", "DBPASS"} {
		if os.Getenv(i) == "" {
			log.Fatalln("Missing env var ", i)
		}
		config[i] = os.Getenv(i)
	}

	// Optional config options default
	optConfig := map[string]string{
		"DBADDR": "localhost",
		"DBPORT": "5432",
		"DBNAME": "handegg",
	}
	// check env vars
	for i := range optConfig {
		switch os.Getenv(i) {
		case "":
			config[i] = optConfig[i]
		default:
			config[i] = os.Getenv(i)
			//log.Tracef("%v: '%v'", i, os.Getenv(i))

		}
	}

	// open DB in gorm
	//var err error
	db, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		config["DBADDR"], config["DBPORT"], config["DBUSER"],
		config["DBNAME"], config["DBPASS"])),
		&gorm.Config{})
	if err != nil {
		log.Fatalln("DB could not open:", err)
	}

	if err := db.AutoMigrate(&userCreds{}); err != nil {
		log.Fatalln("DB failed automigration: ", err)
	}



}
