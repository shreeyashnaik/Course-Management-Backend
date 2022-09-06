package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/shreeyashnaik/Course-Management-Backend/config"

	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {

	try := 0
	for try < 5 {
		if db != nil {
			return db
		}
		db = Connect()
		try++
	}

	return nil
}

func Connect() *gorm.DB {
	sqlDB, err := sql.Open("postgres", config.DB_URI)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB}), &gorm.Config{
		// Turn on logger mode to display corresponding SQL statements w.r.t APIs
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the Database!")
	return db
}
