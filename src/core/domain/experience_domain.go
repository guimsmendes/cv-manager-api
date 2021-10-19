package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ExperienceDomain struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title" validate:"required"`
	Category string `bson:"category" json:"category" validate:"required"`
	Company string `bson:"company" json:"company" validate:"required"`
	Location string `bson:"location" json:"location" validate:"required"`
	Skills []bson.ObjectId `bson:"skills" json:"skills" validate:"required"`
	Description string `bson:"description" json:"description" validate:"required"`
	StartDate time.Time `bson:"startDate" json:"startDate" validate:"required"`
	EndDate time.Time `bson:"endDate" json:"endDate"`
	Logo bson.Binary `bson:"logo" json:"logo"`
}

type ExperienceUseCase interface {
	AddExperience(domain ExperienceDomain) error
	RemoveSkill(id bson.ObjectId) error
	AddSkill(id bson.ObjectId) error
	SetEndDate(time time.Time) error
	ListExperiences() ([]ExperienceDomain, error)
	ListExperiencesByCategory(c string) ([]ExperienceDomain, error)
}
