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
	return s.cvRepository.DbInsert(COLLECTION, domain)
}

func (s *SkillUseCase) RemoveSkill(id string) error {
	return s.cvRepository.DbRemoveOne(COLLECTION, bson.M{"_id": bson.ObjectIdHex(id)})
}

func (s *SkillUseCase) GetSkill(id string) (*domain.SkillDomain, error) {
	doc, err := s.cvRepository.DbFindOne(COLLECTION, bson.M{"_id": bson.ObjectIdHex(id)})
	return s.toSkillDomain(doc), err
}

func (s *SkillUseCase) ListSkills() ([]*domain.SkillDomain, error) {
	docList, err := s.cvRepository.DbFindAll(COLLECTION, bson.M{})
	var skillDomainList []*domain.SkillDomain
	for _, doc := range docList {
		skill := s.toSkillDomain(doc)
		skillDomainList = append(skillDomainList, skill)
	}
	return skillDomainList, err
}

func (s *SkillUseCase) toSkillDomain(doc bson.M) (domain *domain.SkillDomain) {
	bsonBytes, err := bson.Marshal(doc)
	if err != nil {
		return
	}
	err = bson.Unmarshal(bsonBytes, &domain)
	if err != nil {
		return
	}
	return
}

func NewSkillUseCase(repository gateway.CVRepository) domain.SkillUseCase {
	return &SkillUseCase {
		cvRepository: repository,
	}
}




