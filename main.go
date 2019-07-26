package main

import (
	"fmt"
	"gsgo/pkg/setting"
	"gsgo/routers"
	"net/http"

	_ "gsgo/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger API for gsgo
// @version 2.0
// @description This is a sample server Petstore server.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1

func main() {
	router := routers.InitRouter()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
