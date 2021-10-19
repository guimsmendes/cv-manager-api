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
	router.GET("/experience/:keyword", handler.GetExperiencesByKeyword)
	router.GET("/experience", handler.GetAllExperiences)
	router.GET("/experience/:category", handler.GetExperiencesByCategory)
	router.GET("/experience/skill/:id", handler.GetExperiencesBySkillId)
	router.DELETE("/experience/:id/skill/:skillId", handler.DeleteSkill)
}

func (handler *ExperienceHandler) PostExperience(context echo.Context) error {
	experienceDomain := new(domain.ExperienceDomain)

	if err := util.ValidateContext(context, experienceDomain); err != nil {
		return err
	}

	experienceDomain.ID = bson.NewObjectId()

	if err := handler.experienceUseCase.AddExperience(*experienceDomain); err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, experienceDomain)
}

func (handler *ExperienceHandler) GetSkill(context echo.Context) error {
	id := context.Param("id")

	skillDomain, err := handler.skillUseCase.GetSkill(bson.ObjectId(id))

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomain)
}

func (handler *ExperienceHandler) GetAllSkills(context echo.Context) error {
	skillDomainList, err := handler.skillUseCase.ListSkills()

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomainList)
}

func (handler *ExperienceHandler) GetSkillsByCategory(context echo.Context) error {
	category := context.Param("category")

	skillDomain, err := handler.skillUseCase.ListSkillsByCategory(category)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomain)
}

func (handler *ExperienceHandler) DeleteSkill(context echo.Context) error {
	id := context.Param("id")

	err := handler.skillUseCase.RemoveSkill(bson.ObjectId(id))

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, deleted)
}
