package types

import (
	"errors"
	"fmt"
)

type jobList struct {
	list map[string]*Job
}

func NewJobList() *jobList {
	return &jobList{
		list: make(map[string]*Job),
	}
}

var JobList = NewJobList()

func (jl *jobList) ListJobs() []*Job {
	var jobs = make([]*Job, 0, len(jl.list))
	for _, job := range jl.list {
		jobs = append(jobs, job)
	}
	return jobs
}

func (jl *jobList) AddJob(job *Job) {
	jl.list[job.Id] = job
}

func (jl *jobList) GetJob(jobId string) (*Job, error) {
	job, exists := jl.list[jobId]
	if exists {
		return job, nil
	}
	return nil, errors.New("Job with ID '" + jobId + "' does not exist")
}

func (jl *jobList) UpdateJob(jobId string, action string) (*Job, error) {
	job, err := jl.GetJob(jobId)
	if err != nil {
		return nil, err
	}
	fmt.Println("Update this job", job, action)
	return job, nil
}

func (jl *jobList) DeleteJob(jobId string) (*Job, error) {
	job, err := jl.GetJob(jobId)
	if err != nil {
		return nil, errors.New("Job with ID '" + jobId + "' does not exist")
	}
	fmt.Println("Delete this job", job)
	delete(jl.list, jobId)
	return job, nil
}
