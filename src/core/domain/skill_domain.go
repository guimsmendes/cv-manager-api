package domain

import "gopkg.in/mgo.v2/bson"

type SkillDomain struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name" validate:"required"`
	Category string `bson:"category" json:"category" validate:"required"`
	Version float32 `bson:"version" json:"version"`
	Badge bool `bson:"badge" json:"badge"`
	Logo bson.Binary `bson:"logo" json:"logo"`
}

type SkillUseCase interface {
	AddSkill(domain SkillDomain) error
	RemoveSkill(id bson.ObjectId) error
	GetSkill(id bson.ObjectId) (SkillDomain, error)
	ListSkills() ([]SkillDomain, error)
}



