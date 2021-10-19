package main

import (
	"cv-manager-api/src/config/instana"
	"cv-manager-api/src/config/logger"
	"cv-manager-api/src/core/usecase"
	"cv-manager-api/src/dataprovider/mongo"
	"cv-manager-api/src/dataprovider/mongo/repository"
	"cv-manager-api/src/entrypoint/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	router = echo.New()
	instanaInit = instana.NewInstana(logger.NewLogger())
)

func init () {
	
}

func main() {
	router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1<<10,
		LogLevel: log.ERROR,
	}))
	router.Use(middleware.Timeout())
	router.Use(middleware.CORS())

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	var mongoConnection = mongo.Connect("localhost:27017", "guest", "guest","CV")
	var cvRepository = repository.NewCVRepository(mongoConnection)
	var skillUseCase = usecase.NewSkillUseCase(cvRepository)

	rest.NewSkillHandler(router, skillUseCase)
	rest.NewHealthCheckHandler(router, instanaInit)

	var port = "8000"

	logger.Info("Starting server...")
	if err := router.Start(":" + port); err != nil {
		logger.Error("Server has refused to start: ", err)
	}


}