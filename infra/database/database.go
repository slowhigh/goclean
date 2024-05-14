package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/slowhigh/goclean/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(conf *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.DbName,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                             // Slow SQL threshold
			LogLevel:                  logger.LogLevel(conf.Database.LogLevel), // Log level
			IgnoreRecordNotFoundError: true,                                    // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                                   // Don't include params in the SQL log
			Colorful:                  false,                                   // Disable color
		},
	)

	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: postgres,
	}, nil
}
