package rest

import (
	"cv-manager-api/src/core/domain"
	"cv-manager-api/src/entrypoint/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

const (
	deleted = "Skill deleted"
)

type SkillHandler struct {
	skillUseCase domain.SkillUseCase
}

func NewSkillHandler(router *echo.Echo, skillUseCase domain.SkillUseCase) {
	handler := SkillHandler{
		skillUseCase: skillUseCase,
	}

	router.Validator = &util.CustomValidator{
		Validator: validator.New()}

	router.POST("/skill", handler.PostSkill)
	router.GET("/skill/:id", handler.GetSkill)
	router.GET("/skill", handler.GetAllSkills)
	router.DELETE("/skill/:id", handler.DeleteSkill)
}

func (handler *SkillHandler) PostSkill(context echo.Context) error {
	skillDomain := new(domain.SkillDomain)

	if err := util.ValidateContext(context, skillDomain); err != nil {
		return err
	}

	skillDomain.ID = bson.NewObjectId()
	if err := handler.skillUseCase.AddSkill(*skillDomain); err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, skillDomain)
}

func (handler *SkillHandler) GetSkill(context echo.Context) error {
	id := context.Param("id")
	skillDomain, err := handler.skillUseCase.GetSkill(id)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomain)
}

func (handler *SkillHandler) GetAllSkills(context echo.Context) error {
	skillDomainList, err := handler.skillUseCase.ListSkills()

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, skillDomainList)
}

func (handler *SkillHandler) DeleteSkill(context echo.Context) error {
	id := context.Param("id")

	err := handler.skillUseCase.RemoveSkill(id)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, deleted)
}
