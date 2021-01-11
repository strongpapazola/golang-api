package conn

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	// host := os.Getenv("MONGO_HOST")
	// dbName := os.Getenv("MONGO_DB_NAME")
	host := "117.53.47.252"
	dbName := "microservice"
	session, err := mgo.Dial(host)
	log.Println("[*] Connecting To Database ...")
	if err != nil {
		log.Println("[-] Database Error Session :", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
	log.Println("[+] Connected To Database ...")
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
