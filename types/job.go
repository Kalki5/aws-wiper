package types

import (
	"errors"

	"github.com/google/uuid"
)

type JobParam struct {
	Accounts      []Account           `json:"accounts" binding:"required"`
	Regions       []Region            `json:"regions" binding:"required"`
	ResourceTypes map[string][]string `json:"resource_types" binding:"required"`
}

type Job struct {
	Id           string      `json:"id"`
	Status       string      `json:"status"`
	Param        JobParam    `json:"param"`
	Resources    []Resource  `json:"resources"`
	Dependencies [][2]string `json:"dependencies"`
	Relations    [][3]string `json:"relations"`
}

func NewJob(param JobParam) (*Job, error) {
	for _, r := range param.Regions {
		if !IsValidRegion(r) {
			return nil, errors.New("'" + string(r) + "' is not a valid region")
		}
	}

	j := new(Job)
	j.Id = uuid.New().String()
	j.Status = "created"

	j.Param = param

	j.Resources = make([]Resource, 0)
	j.Dependencies = make([][2]string, 0)
	j.Relations = make([][3]string, 0)

	return j, nil
}

func (j *Job) ToMapMini() Map {
	accounts := make([]string, 0, len(j.Param.Accounts))
	for _, a := range j.Param.Accounts {
		accounts = append(accounts, a.Name)
	}

	regions := make([]string, 0, len(j.Param.Regions))
	for _, r := range j.Param.Regions {
		regions = append(regions, r.ToString())
	}

	return Map{
		"id":     j.Id,
		"status": j.Status,
		"param": Map{
			"accounts":       accounts,
			"regions":        regions,
			"resource_types": j.Param.ResourceTypes,
		},
	}
}

func (j *Job) ToMap() Map {
	job := j.ToMapMini()

	resources := make([]Map, 0, len(j.Resources))
	for _, r := range j.Resources {
		resources = append(resources, r.ToMap())
	}

	job["resources"] = resources
	job["dependencies"] = j.Dependencies
	job["relations"] = j.Relations

	return job
}
