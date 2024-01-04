package main

import (
	"clean-arch-go-grpc/pkg/viper"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func main() {
	viper := viper.NewViper()

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.name")

	var db *sql.DB

	// setup database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	log.Printf("sd %s", dsn)
	db, err := sql.Open("postgres", dsn)

	if err = db.Ping(); err != nil {
		log.Fatalln(string("\033[31m"), "error connection: ", err.Error())
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	s, _ := goose.GetDBVersion(db)
	fmt.Println("version of db", s)

	if err := goose.Run("up", db, "./database/migration"); err != nil {
		log.Printf("error %s", err.Error())
	}
}
