package cmd

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	//DBHOST will be stored from corresponding environment
	DBHOST = "DB_HOST"
	//DBPORT will be stored from corresponding environment
	DBPORT = "DB_PORT"
	//DBUSER will be stored from corresponding environment
	DBUSER = "DB_USER"
	//DBPASSWORD will be stored from corresponding environment
	DBPASSWORD = "DB_PASSWORD"
	//DBNAME will be stored from corresponding environment
	DBNAME = "DB_NAME"
)

func getPostgresdb() *sql.DB {
	env := getEnv()
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		env[DBHOST],
		env[DBPORT],
		env[DBUSER],
		env[DBPASSWORD],
		env[DBNAME],
	)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func getEnv() map[string]string {
	host, ok := os.LookupEnv(DBHOST)
	if !ok {
		fmt.Printf("%v HOST environment variable required but not set", DBHOST)
	}
	port, ok := os.LookupEnv(DBPORT)
	if !ok {
		fmt.Printf("%v HOST environment variable required but not set", DBPORT)
	}
	db, ok := os.LookupEnv(DBNAME)
	if !ok {
		fmt.Printf("%v HOST environment variable required but not set", DBNAME)
	}
	psw, ok := os.LookupEnv(DBPASSWORD)
	if !ok {
		fmt.Printf("%v HOST environment variable required but not set", DBPASSWORD)
	}
	user, ok := os.LookupEnv(DBUSER)
	if !ok {
		fmt.Printf("%v HOST environment variable required but not set", DBUSER)
	}
	env := map[string]string{
		DBHOST:     host,
		DBPORT:     port,
		DBNAME:     db,
		DBPASSWORD: psw,
		DBUSER:     user,
	}
	return env
}
