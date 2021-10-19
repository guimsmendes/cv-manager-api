package usecase

import (
	"cv-manager-api/src/core/domain"
	"cv-manager-api/src/core/gateway"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "skill"

type SkillUseCase struct {
	cvRepository gateway.CVRepository
}

func (s *SkillUseCase) AddSkill(domain domain.SkillDomain) error {
	err := s.cvRepository.DbInsert(COLLECTION, domain)
	return err
}


func (s *SkillUseCase) RemoveSkill(id bson.ObjectId) error {
	panic("implement me")
}

func (s *SkillUseCase) GetSkill(id bson.ObjectId) (domain.SkillDomain, error) {
	panic("implement me")
}

func (s *SkillUseCase) ListSkills() ([]domain.SkillDomain, error) {
	panic("implement me")
}

func NewSkillUseCase(repository gateway.CVRepository) domain.SkillUseCase {
	return &SkillUseCase {
		cvRepository: repository,
	}
}




