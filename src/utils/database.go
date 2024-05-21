/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * Manage the connection to the unique MySQL database.
*/

package utils

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DatabaseConnnection *sql.DB // Active connection to the database

func ConnectDB() {
	/* 
		Void func to connect to the MySQL database.
	*/
	
	var err error
	DatabaseConnnection, err = sql.Open("mysql", "root:paul@tcp(127.0.0.1:3306)/shifumi") // TODO: Change to ENV variable
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DatabaseConnnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}