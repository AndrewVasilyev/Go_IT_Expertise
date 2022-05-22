package storage

import (
	"log"
	"os"

	cfg "github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/config"
	yaml "gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {

	//protocol://username:password@host:port/database

	conf := &cfg.DbConfig{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		panic(err)
	}

	dbURL := conf.Protocol + "://" + conf.Username + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.DB

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func NewDB() *gorm.DB {

	return InitDB("db/config.yml")

}
