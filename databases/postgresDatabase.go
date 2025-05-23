package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/bigstth/isekai-shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", conf.DBName)

		postgresDatabaseInstance = &postgresDatabase{conn}
	})

	return postgresDatabaseInstance
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return db.DB
}
