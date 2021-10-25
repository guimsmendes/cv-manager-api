package rest

import (
	"cv-manager-api/src/core/domain"
	"cv-manager-api/src/entrypoint/util"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type ExperienceHandler struct {
	experienceUseCase domain.ExperienceUseCase
}

func NewExperienceHandler(router *echo.Echo, experienceUseCase domain.ExperienceUseCase) {
	handler := ExperienceHandler{
		experienceUseCase: experienceUseCase,
	}

	router.POST("/experience", handler.PostExperience)
	router.POST("/experience/:id/skill/:skillId", handler.PostSkill)
	router.PATCH("/experience/:endDate", handler.PatchEndDate)
	router.GET("/experience", handler.GetAllExperiences)
	router.GET("/experience/:keyword", handler.GetExperiencesByKeyword)
	router.DELETE("/experience/:id/skill/:skillId", handler.DeleteSkill)
}

func (handler *ExperienceHandler) PostExperience(context echo.Context) error {
	experienceDomain := new(domain.ExperienceDomain)

	if err := util.ValidateContext(context, experienceDomain); err != nil {
		return err
	}

	experienceDomain.ID = bson.NewObjectId().String()

	if err := handler.experienceUseCase.AddExperience(*experienceDomain); err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, experienceDomain)
}

func (handler *ExperienceHandler) PostSkill(context echo.Context) error {
	id := context.Param("id")
	skillId := context.Param("skillId")

	experienceDomain, err := handler.experienceUseCase.AddSkill(bson.ObjectId(id), bson.ObjectId(skillId))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusAccepted, experienceDomain)
}

func (handler *ExperienceHandler) PatchEndDate(context echo.Context) error {
	endDate := context.Param("endDate")

	err := handler.experienceUseCase.SetEndDate(endDate)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusAccepted, "")
}

func (handler *ExperienceHandler) GetAllExperiences(context echo.Context) error {
	skillDomainList, err := handler.experienceUseCase.ListExperiences()

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomainList)
}

func (handler *ExperienceHandler) GetExperiencesByKeyword(context echo.Context) error {
	keyword := context.Param("keyword")

	skillDomainList, err := handler.experienceUseCase.ListExperiencesByKeyword(keyword)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomainList)
}

func (handler *ExperienceHandler) DeleteSkill(context echo.Context) error {
	id := context.Param("id")
	skillId := context.Param("skillId")

	experienceDomain, err := handler.experienceUseCase.RemoveSkill(bson.ObjectId(id), bson.ObjectId(skillId))

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, experienceDomain)
}
