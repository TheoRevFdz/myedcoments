package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
