package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:kuse@test123@tcp(158.108.112.52)/parking@charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	query, err := db.Prepare("SELECT ImageBLOB FROM car_images WHERE 1")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
}
