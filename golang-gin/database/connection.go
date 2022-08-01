package database

import (
	"io/ioutil"

	"CMS/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDsn(host string, user string, pw string, db string, port string, ssl string) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pw, db, port, ssl)
	return dsn
}

func MirgrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&domain.News{})
}

func ConnectDatabase(env domain.DotEnv) (*gorm.DB, error) {
	dsn := GetDsn(env.Host, env.User, env.PW, env.Name, env.Port, env.SSL)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	MirgrateDatabase(db)
	return db, nil

}

func ConnectSqlFile(db *gorm.DB) error {
	query, err := ioutil.ReadFile("script/news.sql")
	if err != nil {
		return err
	}
	err = db.Exec(string(query)).Error
	if err != nil {
		return err
	}
	return nil
}
