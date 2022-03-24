package week02

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Ws@Test!@E1#@tcp(10.83.1.164:3306)/configuration")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	return db
}
