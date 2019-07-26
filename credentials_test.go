package gojenkins

import (
	"fmt"
	"testing"
	"time"
)

func TestUsernamePasswordCredentials(t *testing.T) {
	jenkins := NewJenkinsWithTestData()
	err := jenkins.CreateUsernamePasswordCredential("GLOBAL", "test", "test_user", "testing for golang-jenkins module", "password")
	if err != nil {
		t.Errorf("Error %v\n", err)
	}
}

func TestSecretTextCredentials(t *testing.T) {
	jenkins := NewJenkinsWithTestData()
	err := jenkins.CreateCredentialsSecret("GLOBAL", fmt.Sprintf("test-%v", time.Now().UnixNano()), "testing for golang-jenkins module", "123")
	if err != nil {
		t.Errorf("Error %v\n", err)
	}
}
