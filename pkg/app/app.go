package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/store-management-eda/pkg/config"
)

func Init() (*http.Server, error) {
	if err := config.Init(); err != nil {
		return nil, err
	}

	router := gin.Default()

	app := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.Env.PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return app, nil
}
