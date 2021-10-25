package main

import (
	"context"
	"cv-manager-api/src/config/instana"
	"cv-manager-api/src/config/logger"
	"cv-manager-api/src/core/domain"
	"cv-manager-api/src/core/usecase"
	"cv-manager-api/src/dataprovider/elasticsearch"
	"cv-manager-api/src/dataprovider/elasticsearch/client"
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
	ctx := context.Background()

	router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1<<10,
		LogLevel: log.ERROR,
	}))
	router.Use(middleware.Timeout())
	router.Use(middleware.CORS())

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	var elasticClient = elasticsearch.ConnectIndex("experience", ctx, domain.ExperienceMapping)
	var experienceClient = client.NewExperienceClient(elasticClient)
	var experienceUseCase = usecase.NewExperienceUseCase(experienceClient)
	rest.NewExperienceHandler(router, experienceUseCase)

	var mongoConnection = mongo.Connect("localhost:27017", "", "","local")
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