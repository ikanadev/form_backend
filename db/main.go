package db

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

// DBCon global db connection
var DBCon *pg.DB

// Connect funct to connect to database
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "taylor",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "forms",
	}
	db := pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect database \n")
		os.Exit(100)
	}
	log.Printf("Connectio database successfully\n")
	CreateUserTable(db)
	CreateFormEstTable(db)
	CreateStudentTable(db)
	CreateFormProTable(db)
	CreateFormPreTable(db)
	CreateFormDocTable(db)
	CreateFormInsTable(db)
	return db
}
