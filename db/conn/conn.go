package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	engine   = os.Getenv("DB_ENGINE")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
)

// Conn return connection to database
func Conn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(engine, dsn)

	if err != nil {
		log.Fatalf("We can not connect to database for the next reason: %s", err.Error())
	}

	return db
}
