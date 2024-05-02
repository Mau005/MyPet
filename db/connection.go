package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/Mau005/MyPet/configuration"
	"github.com/Mau005/MyPet/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func AutoMigrate() {
	DB.AutoMigrate(&models.Account{})
}

func connectionSqlite(database configuration.DataBase, debugMode bool) error {
	var err error

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}

	DB, err = gorm.Open(sqlite.Open(database.SqlitePath), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	AutoMigrate()
	log.Println("Sqlite Conection OK")
	return nil
}

func connectionMysql(database configuration.DataBase, debugMode bool) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.User, database.Password, database.Host, database.Port, database.NameDB)

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	AutoMigrate()
	log.Println("Mysql Conection OK")
	return nil
}

func ConnectionDataBase() error {
	switch configuration.Config.DataBase.Engine {
	case "sqlite":
		return connectionSqlite(configuration.Config.DataBase, configuration.Config.Server.Debug)

	case "mysql":
		return connectionMysql(configuration.Config.DataBase, configuration.Config.Server.Debug)

	default:
		return errors.New("no known database detected")
	}

}
