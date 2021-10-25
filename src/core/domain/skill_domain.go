package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type SkillDomain struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name" validate:"required"`
	Category string `json:"category" bson:"category" validate:"required"`
	Version float32 `json:"version" bson:"version"`
	Badge bool `json:"badge" bson:"badge"`
	Logo bson.Binary `json:"logo" bson:"logo"`
}

type SkillUseCase interface {
	AddSkill(domain SkillDomain) error
	RemoveSkill(id string) error
	GetSkill(id string) (*SkillDomain, error)
	ListSkills() ([]*SkillDomain, error)
}



