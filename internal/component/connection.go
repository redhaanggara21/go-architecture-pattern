package component

import (
	"database/sql"
	"fmt"
	"log"

	"red21.id/learn/bengkel/internal/config"
)

func GetConnection(conf config.Config) *sql.DB {
	dsn := fmt.Sprint(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error when connect to database: %s", err.Error())
	}
	return db
}
