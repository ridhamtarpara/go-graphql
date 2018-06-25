package graph

import (
	"context"
	"fmt"
	"log"
	"time"

	"ridham.me/jobs/graph/database"
)

// TODO manage this app object effectively
// App
type App struct {
	jobs []Job
}

// Mutation_createJob create new job in firebase db
func (a *App) Mutation_createJob(ctx context.Context, job NewJob) (Job, error) {
	// Push the reference to get the new key
	jobRef, err := database.DB.NewRef("/Jobs").Push(database.Context, nil)
	if err != nil {
		log.Fatalln("Error reading from database: ", err)
	}
	jobId := jobRef.Key

	// Create job object from request
	// TODO change createdBy
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

	// Set the values in the DB
	if err = database.DB.NewRef("/Jobs").Child(jobId).Set(database.Context, newJob); err != nil {
		fmt.Printf("Database error: ", err)
	}
	log.Println(121212, newJob)

	return newJob, nil
}

// Query_jobs get all jobs
func (a *App) Query_jobs(ctx context.Context) ([]Job, error) {
	var result map[string]Job

	if err := database.DB.NewRef("/Jobs").Get(database.Context, &result); err != nil {
		fmt.Printf("Database error", err)
	}
	var jobs []Job
	// Convert map to array of jobs
	for key, job := range result {
		job.ID = key
		jobs = append(jobs, job)
	}
	return jobs, nil
}
