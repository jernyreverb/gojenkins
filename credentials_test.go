package gojenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	jenkinsPassword = "admin"
)

func TestInvalidCreateCredentials(t *testing.T) {
	_, err := NewCredential("test")
	assert.NotNil(t, err)
}

func TestValidCreateCredentials(t *testing.T) {
	cred, err := NewCredential(CredentialUsernamePassword)

	assert.Nil(t, err)
	assert.NotNil(t, cred)

	cred, err = NewCredential(CredentialSSHKey)

	assert.Nil(t, err)
	assert.NotNil(t, cred)

	cred, err = NewCredential(CredentialSecretText)

	assert.Nil(t, err)
	assert.NotNil(t, cred)
}

func TestUsernamePasswordCredentialsCreation(t *testing.T) {
	cred, err := NewCredential(CredentialUsernamePassword)

	assert.Nil(t, err)
	assert.NotNil(t, cred)

	userCreds, ok := cred.(*CredentialUserNamePayload)
	assert.Equal(t, ok, true)
	userCreds.SetId("test")
	userCreds.SetUsername("test")
	userCreds.SetPassword("test")
	userCreds.SetName("test")
	userCreds.SetDescription("test")

	jenkins := CreateJenkins(nil, "http://localhost:8080", "admin", jenkinsPassword)
	_, err = jenkins.Init()
	assert.Nil(t, err, "Jenkins Initialization should not fail")

	err = jenkins.CreateCredential(cred)

	assert.Nil(t, err)
}

func TestSecretTextCredentialsCreation(t *testing.T) {
	cred, err := NewCredential(CredentialSecretText)

	assert.Nil(t, err)
	assert.NotNil(t, cred)

	userCreds, ok := cred.(*CredentialSecretTextPayload)
	assert.Equal(t, ok, true)
	userCreds.SetId("test-secret")
	userCreds.SetName("test")
	userCreds.SetDescription("test")
	userCreds.SetSecret("wow")

	jenkins := CreateJenkins(nil, "http://localhost:8080", "admin", jenkinsPassword)
	_, err = jenkins.Init()
	assert.Nil(t, err, "Jenkins Initialization should not fail")

	err = jenkins.CreateCredential(cred)

	assert.Nil(t, err)
}

func TestSSHKeyCredentialsCreation(t *testing.T) {
	cred, err := NewCredential(CredentialSSHKey)

	assert.Nil(t, err)
	assert.NotNil(t, cred)

	userCreds, ok := cred.(*CredentialSSHKeyPayload)
	assert.Equal(t, ok, true)
	userCreds.SetId("test-secret-ssh")
	userCreds.SetName("test")
	userCreds.SetDescription("test")
	userCreds.SetPrivateKey("adfadfafafadfadfadfasdfadfadfafafd")

	jenkins := CreateJenkins(nil, "http://localhost:8080", "admin", jenkinsPassword)
	_, err = jenkins.Init()
	assert.Nil(t, err, "Jenkins Initialization should not fail")

	err = jenkins.CreateCredential(cred)

	assert.Nil(t, err)
}
