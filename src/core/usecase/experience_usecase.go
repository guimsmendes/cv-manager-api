package usecase

import (
	"cv-manager-api/src/core/domain"
	"cv-manager-api/src/core/gateway"
	"gopkg.in/mgo.v2/bson"
)

type ExperienceUseCase struct {
	client gateway.ExperienceClient
}

func (e ExperienceUseCase) AddExperience(domain domain.ExperienceDomain) error {
	panic("implement me")
}

func (e ExperienceUseCase) RemoveSkill(id bson.ObjectId, skillId bson.ObjectId) (domain.ExperienceDomain, error) {
	panic("implement me")
}

func (e ExperienceUseCase) AddSkill(id bson.ObjectId, skillId bson.ObjectId) (domain.ExperienceDomain, error) {
	panic("implement me")
}

func (e ExperienceUseCase) SetEndDate(endDate string) error {
	panic("implement me")
}

func (e ExperienceUseCase) ListExperiences() ([]domain.ExperienceDomain, error) {
	panic("implement me")
}

func (e ExperienceUseCase) ListExperiencesByKeyword(keyword string) ([]domain.ExperienceDomain, error) {
	panic("implement me")
}

func NewExperienceUseCase(client gateway.ExperienceClient) domain.ExperienceUseCase {
	return &ExperienceUseCase {
		client: client,
	}
}