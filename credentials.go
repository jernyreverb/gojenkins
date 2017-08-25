package gojenkins

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

const CredentialUsernamePassword = "usernamepassword"
const CredentialSSHKey = "sshkey"
const CredentialSecretText = "secrettext"

type Credential interface {
	Payload() io.Reader
}

type PayloadWrapper struct {
	Id         string     `json:"_"`
	Credential Credential `json:"credentials"`
}

func NewCredential(credType string) (Credential, error) {
	switch credType {
	case CredentialUsernamePassword:
		return &CredentialUserNamePayload{
			Scope: "GLOBAL",
			Class: "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl",
		}, nil
	case CredentialSSHKey:
		return &CredentialSSHKeyPayload{
			Scope: "GLOBAL",
			Class: "com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey",
			PrivateKeySource: &PrivateKeySource{
				Class: "com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey$DirectEntryPrivateKeySource",
			},
		}, nil
	case CredentialSecretText:
		return &CredentialSecretTextPayload{
			Scope: "GLOBAL",
			Class: "org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl",
		}, nil
	default:
		return nil, errors.New("invalid Credential type")
	}
}

type CredentialUserNamePayload struct {
	Scope       string `json:"scope"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Class       string `json:"$class"`
}

func (c *CredentialUserNamePayload) Payload() io.Reader {
	str, _ := json.Marshal(&PayloadWrapper{
		Id:         "0",
		Credential: c,
	})
	return strings.NewReader("json=" + string(str))
}

func (c *CredentialUserNamePayload) SetId(id string) {
	c.Id = id
}

func (c *CredentialUserNamePayload) SetDescription(description string) {
	c.Description = description
}

func (c *CredentialUserNamePayload) SetName(name string) {
	c.Name = name
}

func (c *CredentialUserNamePayload) SetUsername(name string) {
	c.Username = name
}

func (c *CredentialUserNamePayload) SetPassword(password string) {
	c.Password = password
}

type CredentialSSHKeyPayload struct {
	Scope            string            `json:"scope"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	Username         string            `json:"username"`
	PrivateKeySource *PrivateKeySource `json:"privateKeySource"`
	Description      string            `json:"description"`
	Class            string            `json:"$class"`
}

type PrivateKeySource struct {
	PrivateKey string `json:"privateKey"`
	Class      string `json:"$class"`
}

func (c *CredentialSSHKeyPayload) Payload() io.Reader {
	str, _ := json.Marshal(&PayloadWrapper{
		Id:         "0",
		Credential: c,
	})
	return strings.NewReader("json=" + string(str))
}

func (c *CredentialSSHKeyPayload) SetId(id string) {
	c.Id = id
}

func (c *CredentialSSHKeyPayload) SetDescription(description string) {
	c.Description = description
}

func (c *CredentialSSHKeyPayload) SetName(name string) {
	c.Name = name
}

func (c *CredentialSSHKeyPayload) SetUsername(name string) {
	c.Username = name
}

func (c *CredentialSSHKeyPayload) SetPrivateKey(privateKey string) {
	c.PrivateKeySource.PrivateKey = privateKey
}

type CredentialSecretTextPayload struct {
	Scope       string `json:"scope"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	Description string `json:"description"`
	Class       string `json:"$class"`
}

func (c *CredentialSecretTextPayload) Payload() io.Reader {
	str, _ := json.Marshal(&PayloadWrapper{
		Id:         "0",
		Credential: c,
	})
	return strings.NewReader("json=" + string(str))
}

func (c *CredentialSecretTextPayload) SetId(id string) {
	c.Id = id
}

func (c *CredentialSecretTextPayload) SetDescription(description string) {
	c.Description = description
}

func (c *CredentialSecretTextPayload) SetName(name string) {
	c.Name = name
}

func (c *CredentialSecretTextPayload) SetSecret(secret string) {
	c.Secret = secret
}
