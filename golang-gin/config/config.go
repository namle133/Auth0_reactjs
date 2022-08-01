package config

import (
	"fmt"
	"os"

	"CMS/database"
	"CMS/domain"
	"CMS/log"
	"CMS/service"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var DotE domain.DotEnv
var UserService service.UserService
var Conf *config.Config

func InitDB() {
	DB, err := database.ConnectDatabase(DotE)
	if err != nil {
		log.Logrus().Error(err)
		return
	}
	if DB == nil {
		logrus.Panic("Could not connect database")
	}
	us := service.UserService{Db: DB}
	UserService = us
	err = database.ConnectSqlFile(DB)
	if err != nil {
		log.Logrus().Error(err)
		return
	}
}
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file": ".env",
		}).Error("Failed to load file")
	}
	host := GetEnv("DB_HOST", "127.0.0.1")

	user := GetEnv("DB_USER", "")
	pw := GetEnv("DB_PASSWORD", "")
	name := GetEnv("DB_NAME", "")
	port := GetEnv("DB_PORT", "5432")
	ssl := GetEnv("DB_SSLMODE", "require")
	err = CheckDatabaseEnv(host, user, pw, name, port, ssl)
	if err != nil {
		log.Logrus().Error(err)
		return
	}
	DotE = domain.DotEnv{Host: host, User: user, PW: pw, Name: name, Port: port, SSL: ssl}

}

func CheckDatabaseEnv(host string, user string, pw string, name string, port string, ssl string) error {
	if len(host) == 0 {
		return fmt.Errorf("%s is empty", host)
	} else if len(user) == 0 {
		return fmt.Errorf("%s is empty", user)
	} else if len(pw) == 0 {
		return fmt.Errorf("%s is empty", pw)
	} else if len(name) == 0 {
		return fmt.Errorf("%s is empty", name)
	} else if len(port) == 0 {
		return fmt.Errorf("%s is empty", port)
	} else if len(ssl) == 0 {
		return fmt.Errorf("%s is empty", ssl)
	}
	return nil

}

func GetEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
