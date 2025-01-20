package db

import (
	"fmt"

	"craftnet/config"
	"craftnet/internal/util"

	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
)

var Instance *sql.DB

// connect to DB
func ConnectDatabase() {
	// read config data from AppConfig
	dbConfig := config.AppConfig.Database

	// Capture connection properties.
	cfg := mysql.Config{
		User:   dbConfig.User,
		Passwd: dbConfig.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		DBName: dbConfig.DBName,
		AllowNativePasswords: dbConfig.AllowNativePasswords,
	}

	// Get a database handle.
	var err error
	Instance, err = sql.Open("mysql", cfg.FormatDSN())
	if !lo.IsNil(err) {
		util.GetLogger().LogErrorWithMsgAndError("Unable to connect to database", err, true)
	}

	pingErr := Instance.Ping()
	if pingErr != nil {
		util.GetLogger().LogErrorWithMsgAndError("Unable to connect to database", pingErr, true)
	}

	util.GetLogger().LogInfo("Connected succesfully to database")
}
