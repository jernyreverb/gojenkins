package gojenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNodes(t *testing.T) {
	t.Skip("not running create nodes")

	id1 := "node1_test"
	id2 := "node2_test"
	id3 := "node3_test"

	jnlp := map[string]string{"method": "JNLPLauncher"}
	ssh := map[string]string{"method": "SSHLauncher"}

	node1, _ := jenkins.CreateNode(id1, 1, "Node 1 Description", "/var/lib/jenkins", "", jnlp)
	if node1 == nil {
		t.Fatal("response should not be nil")
	}

	assert.Equal(t, id1, node1.GetName())

	node2, _ := jenkins.CreateNode(id2, 1, "Node 2 Description", "/var/lib/jenkins", "jdk8 docker", ssh)

	if node2 == nil {
		t.Fatal("response should not be nil")
	}

	assert.Equal(t, id2, node2.GetName())

	node3, _ := jenkins.CreateNode(id3, 1, "Node 3 Description", "/var/lib/jenkins", "jdk7")

	if node3 == nil {
		t.Fatal("response should not be nil")
	}
	assert.Equal(t, id3, node3.GetName())
}

func TestGetAllNodes(t *testing.T) {
	t.Skip("skipping node creationg")
	nodes, _ := jenkins.GetAllNodes()
	assert.Equal(t, 4, len(nodes))
	assert.Equal(t, nodes[0].GetName(), "master")
}

func TestGetLabel(t *testing.T) {
	t.Skip("skipping labels")
	label, err := jenkins.GetLabel("test_label")
	assert.Nil(t, err)
	assert.Equal(t, label.GetName(), "test_label")
	assert.Equal(t, 0, len(label.GetNodes()))

	label, err = jenkins.GetLabel("jdk7")
	assert.Nil(t, err)
	assert.Equal(t, label.GetName(), "jdk7")
	assert.Equal(t, 1, len(label.GetNodes()))
	assert.Equal(t, "node3_test", label.GetNodes()[0].NodeName)

	label, err = jenkins.GetLabel("jdk8")
	assert.Nil(t, err)
	assert.Equal(t, label.GetName(), "jdk8")
	assert.Equal(t, 1, len(label.GetNodes()))
	assert.Equal(t, "node2_test", label.GetNodes()[0].NodeName)

	label, err = jenkins.GetLabel("docker")
	assert.Nil(t, err)
	assert.Equal(t, label.GetName(), "docker")
	assert.Equal(t, 1, len(label.GetNodes()))
	assert.Equal(t, "node2_test", label.GetNodes()[0].NodeName)
}
