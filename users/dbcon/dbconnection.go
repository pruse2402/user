package dbcon

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"

	"user/users/conf"
)

var db *pg.DB

//Connect database
func Connect() {
	dbCon := pg.Connect(&pg.Options{
		Addr:     conf.DatabaseAddr,
		User:     conf.DatabaseUsername,
		Password: conf.DatabasePassword,
		Database: conf.DatabaseName,
	})

	db = dbCon
	log.Printf("Connected successfully")

	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
		log.Fatalf("Unable to connect Postgres Database: %v\n", err)
		os.Exit(1)
	}
}

//Get db connection
func Get() *pg.DB {
	return db
}

//Close db connection
func Close() {
	err := db.Close()

	if err != nil {
		log.Printf("Closing DB err", err)
	}
	log.Printf("DB closed")
}
