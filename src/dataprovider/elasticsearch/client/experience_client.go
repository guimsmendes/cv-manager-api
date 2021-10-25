package client

import (
	"cv-manager-api/src/core/gateway"
	"github.com/olivere/elastic"
)

type ExperienceClient struct { *elastic.Client }


func NewExperienceClient(client *elastic.Client) gateway.ExperienceClient {
	return &ExperienceClient{client}
}