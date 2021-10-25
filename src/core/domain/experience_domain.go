package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ExperienceDomain struct {
	ID string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
	Category string `json:"category" validate:"required"`
	Company string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Skills []string `json:"skills" validate:"required"`
	Description string `json:"description" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate time.Time `json:"endDate"`
	Logo string `json:"logo"`
}

type ExperienceUseCase interface {
	AddExperience(domain ExperienceDomain) error
	RemoveSkill(id bson.ObjectId, skillId bson.ObjectId) (ExperienceDomain, error)
	AddSkill(id bson.ObjectId, skillId bson.ObjectId)  (ExperienceDomain, error)
	SetEndDate(endDate string) error
	ListExperiences() ([]ExperienceDomain, error)
	ListExperiencesByKeyword(keyword string) ([]ExperienceDomain, error)
}

const ExperienceMapping = `
{
	"settings":{
		"number_of_shards": 3,
		"number_of_replicas": 0
	},
	"mappings":{
			"properties":{
				"id":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				},
				"title":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				},
				"category":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				},
				"company":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				},
				"location":{
					"type":"geo_point",
					"index": true,
                    "analyzer": "portuguese"
				},
				"skills":{
					"type":"keyword",
					"index": true,
                    "analyzer": "portuguese"
				},
				"description":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				},
				"startDate":{
					"type":"date",
					"index": true,
                    "analyzer": "portuguese"
				},
				"endDate":{
					"type":"date",
					"index": true,
                    "analyzer": "portuguese"
				},
				"logo":{
					"type":"text",
					"index": true,
                    "analyzer": "portuguese"
				}
			}
	}
}`