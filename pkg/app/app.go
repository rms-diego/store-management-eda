package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/store-management-eda/pkg/config"
	"github.com/rms-diego/store-management-eda/pkg/database"
)

func Init() (*http.Server, *gin.Engine, error) {
	if err := config.Init(); err != nil {
		return nil, nil, err
	}

	if err := database.Init(config.Env.DATABASE_URL); err != nil {
		panic(err)
	}

	router := gin.Default()

	app := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.Env.PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running", "docs": "/docs"})
	})

	return app, router, nil
}
