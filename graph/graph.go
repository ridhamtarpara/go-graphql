package graph

import (
	"context"
	"fmt"
	"time"
	"log"

	"ridham.me/jobs/graph/database"
)

// App ...
type App struct {
	jobs []Job
}

// Mutation_createJL̥ob ...
func (a *App) Mutation_createJob(ctx context.Context, job NewJob) (Job, error) {
	jobRef, err := database.DB.NewRef("/Jobs").Push(database.Context, nil)
	if err != nil {
		log.Fatalln("Error reading from database: ", err)
	}
	jobId := jobRef.Key

	newJob := Job{
		ID:           jobId,
		Name:         job.Name,
		Country:      job.Country,
		Description:  job.Description,
		NoOfPosition: job.NoOfPosition,
		IsDeleted:    job.IsDeleted,
		CreatedBy:    "Admin",
		CreatedAt:    fmt.Sprint(time.Now().Format("2006.01.02 15:04:05")),
	}
	fmt.Printf("New Job : ", newJob)
	//
	if err = database.DB.NewRef("/Jobs").Child(jobId).Set(database.Context, newJob); err != nil {
		fmt.Printf("Database error: ", err)
	}
	log.Println(121212, newJob)

	return newJob, nil
}

func (a *App) Query_jobs(ctx context.Context) ([]Job, error) {
	var result map[string]Job

	if err := database.DB.NewRef("/Jobs").Get(database.Context, &result); err != nil {
		fmt.Printf("Database error", err)
	}
	var jobs []Job
	for key, job := range result {
		job.ID = key
		jobs = append(jobs, job)
	}
	return jobs, nil
}
