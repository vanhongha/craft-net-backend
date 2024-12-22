package db

import (
	"context"
	"fmt"
	"time"

	"craftnet/config"
	"craftnet/internal/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func GetDB() *pgxpool.Pool {
	if DB == nil {
		ConnectDatabase()
	}
	if DB == nil {
		util.GetLogger().LogErrorWithMsg("Please connect to DB first", true)
	}
	return DB
}

// connect to DB
func ConnectDatabase() {
	// read config data from AppConfig
	dbConfig := config.AppConfig.Database

	// connection string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.SSLMode,
	)

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Unable to connect to database", err, false)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.Ping(ctx); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Unable to ping database", err, false)
	}

	util.GetLogger().LogInfo("Connected succesfully to PostgreSQL")
}

// close DB
func CloseDatabase() {
	DB.Close()
	util.GetLogger().LogInfo("Database connection closed")
}
