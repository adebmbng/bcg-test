package main

import (
	"fmt"
	"github.com/adebmbng/bcg-test/configs"
	"github.com/adebmbng/bcg-test/domains/inventories"
	"github.com/adebmbng/bcg-test/domains/promos"
	mysql2 "github.com/adebmbng/bcg-test/repositories/mysql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// main to run entire project (specially graphql server)
func main() {
	// get envar
	cfg := configs.Get()

	// init repository
	mysqlDB := initDB(cfg)
	mysqlRepo := mysql2.NewRepository(mysqlDB)

	// init domain service
	promo := promos.NewPromos(mysqlRepo)
	inventory := inventories.NewInventories(mysqlRepo, promo)

	// TODO: Create GraphQL server
}

// initDB to init mysql repo
func initDB(cfg configs.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	/// local mode development
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		cfg.DBUserName, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}
	// db.DB().SetConnMaxLifetime(time.Duration(int(time.Minute) * cfg.DBMaxLifetimeConnection))

	return db
}
