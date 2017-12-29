package gojenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateViews(t *testing.T) {
	setupJenkins(t)

	list_view, err := jenkins.CreateView("test_list_view", LIST_VIEW)
	assert.Nil(t, err)
	require.NotNil(t, list_view)
	assert.Equal(t, "test_list_view", list_view.GetName())
	assert.Equal(t, "", list_view.GetDescription())
	assert.Equal(t, 0, len(list_view.GetJobs()))

	my_view, err := jenkins.CreateView("test_my_view", MY_VIEW)
	assert.Nil(t, err)
	assert.Equal(t, "test_my_view", my_view.GetName())
	assert.Equal(t, "", my_view.GetDescription())
	assert.Equal(t, 2, len(my_view.GetJobs()))

}

func TestGetViews(t *testing.T) {
	setupJenkins(t)

	views, _ := jenkins.GetAllViews()
	assert.Equal(t, len(views), 3)
	assert.Equal(t, len(views[0].Raw.Jobs), 2)
}

func TestGetSingleView(t *testing.T) {
	setupJenkins(t)

	view, _ := jenkins.GetView("All")
	view2, _ := jenkins.GetView("test_list_view")
	assert.Equal(t, len(view.Raw.Jobs), 2)
	assert.Equal(t, len(view2.Raw.Jobs), 0)
	assert.Equal(t, view2.Raw.Name, "test_list_view")
}
