package gojenkins

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildIsQueued(t *testing.T) {
	setupJenkins(t)

	jobConfig := &jobConfig{
		Name: "job_queued_test",
		File: "job_queued.xml",
	}

	job, err := getOrCreateJob(jobConfig)

	require.Nil(t, err)
	require.NotNil(t, job)

	job.InvokeSimple(map[string]string{"param1": "param1"})
	job.Poll()

	isQueued, err := job.IsQueued()
	require.Nil(t, err)
	require.Equal(t, true, isQueued)

	cleanupJobQueue(jobConfig.Name, t)
}

func TestBuildCreate(t *testing.T) {
	setupJenkins(t)

	jobConfig := &jobConfig{
		Name: "Job1_test",
		File: "job.xml",
	}

	job, _ := getOrCreateJob(jobConfig)

	require.NotNil(t, job)
	job.InvokeSimple(map[string]string{"param1": "param1"})
	job.Poll()

	time.Sleep(10 * time.Second)
	builds, _ := job.GetAllBuildIds()
	assert.True(t, (len(builds) > 0))
}

func TestBuildParseHistory(t *testing.T) {
	setupJenkins(t)

	r, err := os.Open("_tests/build_history.txt")
	if err != nil {
		panic(err)
	}
	history := parseBuildHistory(r)
	assert.True(t, len(history) == 3)
}

func TestBuildGetAll(t *testing.T) {
	setupJenkins(t)

	jobConfig := &jobConfig{
		Name: "Job1_test",
		File: "job.xml",
	}

	job, _ := getOrCreateJob(jobConfig)

	require.NotNil(t, job)

	builds, _ := jenkins.GetAllBuildIds(jobConfig.Name)
	for _, b := range builds {
		build, _ := jenkins.GetBuild(jobConfig.Name, b.Number)
		assert.Equal(t, "SUCCESS", build.GetResult())
	}
	assert.Equal(t, 1, len(builds))
}

func TestBuildMethods(t *testing.T) {
	setupJenkins(t)

	jobConfig := &jobConfig{
		Name: "Job1_test",
		File: "job.xml",
	}

	job, _ := getOrCreateJob(jobConfig)
	require.NotNil(t, job)
	build, _ := job.GetLastBuild()
	require.NotNil(t, build)
	params := build.GetParameters()
	require.NotNil(t, params)
	assert.Equal(t, "params1", params[0].Name)
}
