package gojenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateJobs(t *testing.T) {
	setupJenkins(t)

	job1ID := "Job1_test"
	job2ID := "job2_test"
	job_data := getFileAsString("job.xml")

	job1, err := jenkins.CreateJob(job_data, job1ID)
	assert.Nil(t, err)
	if assert.NotNil(t, job1) {
		assert.Equal(t, "Some Job Description", job1.GetDescription())
		assert.Equal(t, job1ID, job1.GetName())
	}

	job2, _ := jenkins.CreateJob(job_data, job2ID)
	if assert.NotNil(t, job2) {
		assert.Equal(t, "Some Job Description", job2.GetDescription())
		assert.Equal(t, job2ID, job2.GetName())
	}
}

func TestGetAllJobs(t *testing.T) {
	setupJenkins(t)

	jobs, _ := jenkins.GetAllJobs()
	if assert.NotNil(t, jobs) {
		assert.Equal(t, 2, len(jobs))
		assert.Equal(t, jobs[0].Raw.Color, "blue")
	}
}

func TestGetSingleJob(t *testing.T) {
	setupJenkins(t)

	job, _ := jenkins.GetJob("Job1_test")
	isRunning, _ := job.IsRunning()
	config, err := job.GetConfig()
	assert.Nil(t, err)
	assert.Equal(t, false, isRunning)
	assert.Contains(t, config, "<project>")
}

func TestEnableDisableJob(t *testing.T) {
	setupJenkins(t)

	job, _ := jenkins.GetJob("Job1_test")
	result, _ := job.Disable()
	assert.Equal(t, true, result)
	result, _ = job.Enable()
	assert.Equal(t, true, result)
}

func TestCopyDeleteJob(t *testing.T) {
	setupJenkins(t)

	job, _ := jenkins.GetJob("Job1_test")
	require.NotNil(t, job)
	jobCopy, _ := job.Copy("Job1_test_copy")
	require.NotNil(t, jobCopy)
	assert.Equal(t, jobCopy.GetName(), "Job1_test_copy")
	jobDelete, _ := job.Delete()
	assert.Equal(t, true, jobDelete)
}

func TestCreateJobInFolder(t *testing.T) {
	setupJenkins(t)

	jobName := "Job_test"
	job_data := getFileAsString("job.xml")

	job1, err := jenkins.CreateJobInFolder(job_data, jobName, "folder1_test")
	assert.Nil(t, err)
	require.NotNil(t, job1)
	assert.Equal(t, "Some Job Description", job1.GetDescription())
	assert.Equal(t, jobName, job1.GetName())

	job2, err := jenkins.CreateJobInFolder(job_data, jobName, "folder1_test", "folder2_test")
	assert.Nil(t, err)
	require.NotNil(t, job2)
	assert.Equal(t, "Some Job Description", job2.GetDescription())
	assert.Equal(t, jobName, job2.GetName())
}
