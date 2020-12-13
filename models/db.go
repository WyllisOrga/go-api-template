package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	// usefull to import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/wyllisMonteiro/go-api-template/services"
)

func getVarsDB() (string, string, string, string, error) {
	userDB, err := services.GoDotEnvVariable("USERDB")
	if err != nil {

		return "", "", "", "", err
	}

	passDB, err := services.GoDotEnvVariable("PASSDB")
	if err != nil {
		return "", "", "", "", err
	}

	hostDB, err := services.GoDotEnvVariable("HOSTDB")
	if err != nil {
		return "", "", "", "", err
	}

	nameDB, err := services.GoDotEnvVariable("NAMEDB")
	if err != nil {
		return "", "", "", "", err
	}

	return userDB, passDB, hostDB, nameDB, nil
}

// DB database
var DB *sql.DB

const (
	attemptsDBConnexion = 3
	waitForConnexion    = 3
)

// ConnectToDB Make connexion with database
func ConnectToDB() error {
	userDB, passDB, hostDB, nameDB, err := getVarsDB()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", userDB, passDB, hostDB, nameDB)
	DB, err = sql.Open("mysql", dsn)

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
