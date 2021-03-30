package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	env "github.com/caarlos0/env/v6"
	// usefull to import mysql
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	User string `env:"MYSQL_USER"`
	Pass string `env:"MYSQL_PASSWORD"`
	Database string `env:"MYSQL_DATABASE"`
	Host string `env:"DB_HOST"`
}

//DB database
var DB *sql.DB

const (
	attemptsDBConnexion = 3
	waitForConnexion    = 3
)

//ConnectToDB Make connexion with database
func ConnectToDB() error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Database)
	DB, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	for index := 1; index <= attemptsDBConnexion; index++ {
		err = DB.Ping()
		if err != nil {
			if index < attemptsDBConnexion {
				log.Printf("db connection failed, %d retry : %v", index, err)
				time.Sleep(waitForConnexion * time.Second)
			}
			continue
		}

		break
	}

	if err != nil {
		return errors.New("Can't connect to database after 3 attempts")
	}

	return nil
}
