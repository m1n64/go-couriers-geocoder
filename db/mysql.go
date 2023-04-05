package db

import (
	"awesomeProject/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	user := utils.GoDotEnvVariable("MYSQL_USER")
	password := utils.GoDotEnvVariable("MYSQL_PASSWORD")
	host := utils.GoDotEnvVariable("MYSQL_HOST")
	port := utils.GoDotEnvVariable("MYSQL_PORT")
	database := utils.GoDotEnvVariable("MYSQL_DB")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	return db
}
