package domain

type category string

const (
	BACK_END category = "Back-End"
	FRONT_END category = "Front-End"
	DATA_SCIENCE category = "Data Science"
	DEVOPS category = "DevOps"
	BUSINESS category = "Business"
	PROJECT_MANAGEMENT category = "Project Management"
)

type Categorize interface {
	Category() category
}

func(c category) Category() category {
	return c
}