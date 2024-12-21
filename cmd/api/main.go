package main

import (
	"net/http"

	"craftnet/config"
	"craftnet/internal/db"
	"craftnet/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// init logger
	util.InitLogger("../../logs/app.log")
	defer util.GetLogger().Close()

	// load config file
	config.LoadConfig()

	// connect to DB
	db.ConnectDatabase()
	defer db.CloseDatabase()

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
