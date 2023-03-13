package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

// DBCon global db connection
var DBCon *pg.DB

// Connect funct to connect to database
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
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
