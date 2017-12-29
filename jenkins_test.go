package gojenkins

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	jenkins *Jenkins
)

type jobConfig struct {
	Name string
	File string
}

func setupJenkins(t *testing.T) {
	jenkinsPassword := "admin"

	jenkins = CreateJenkins(nil, "http://localhost:8080", "admin", jenkinsPassword)
	_, err := jenkins.Init()
	assert.Nil(t, err, "Jenkins Initialization should not fail")
}

func getOrCreateJob(config *jobConfig) (*Job, error) {

	job_data := getFileAsString(config.File)
	job, err := jenkins.CreateJob(job_data, config.Name)

	errMsg := "A job already exists with the name  " + config.Name

	if err != nil && err.Error() == errMsg {
		job, _ := jenkins.GetJob(config.Name)
		return job, nil
	}

	return job, nil
}

func cleanupJobQueue(name string, t *testing.T) {
	queue, _ := jenkins.GetQueue()
	require.NotNil(t, queue)
	tasks := queue.GetTasksForJob(name)
	require.NotNil(t, tasks)
	require.NotEmpty(t, tasks)
	queue.CancelTask(tasks[0].Raw.ID)
}

func TestGetPlugins(t *testing.T) {
	setupJenkins(t)

	plugins, _ := jenkins.GetPlugins(3)
	assert.Equal(t, 5, plugins.Count())
}

func TestCreateFolder(t *testing.T) {
	setupJenkins(t)

	folder1ID := "folder1_test"
	folder2ID := "folder2_test"

	folder1, err := jenkins.CreateFolder(folder1ID)
	assert.Nil(t, err)
	require.NotNil(t, folder1)
	assert.Equal(t, folder1ID, folder1.GetName())

	folder2, err := jenkins.CreateFolder(folder2ID, folder1ID)
	assert.Nil(t, err)
	assert.NotNil(t, folder2)
	assert.Equal(t, folder2ID, folder2.GetName())
}

func TestGetFolder(t *testing.T) {
	setupJenkins(t)

	folder1ID := "folder1_test"
	folder2ID := "folder2_test"

	folder1, err := jenkins.GetFolder(folder1ID)
	assert.Nil(t, err)
	assert.NotNil(t, folder1)
	assert.Equal(t, folder1ID, folder1.GetName())

	folder2, err := jenkins.GetFolder(folder2ID, folder1ID)
	assert.Nil(t, err)
	assert.NotNil(t, folder2)
	assert.Equal(t, folder2ID, folder2.GetName())
}

func TestConcurrentRequests(t *testing.T) {
	for i := 0; i <= 16; i++ {
		go func() {
			jenkins.GetAllJobs()
			jenkins.GetAllViews()
			jenkins.GetAllNodes()
		}()
	}
}

func getFileAsString(path string) string {
	buf, err := ioutil.ReadFile("_tests/" + path)
	if err != nil {
		panic(err)
	}

	return string(buf)
}
