package gojenkins

import (
	"testing"
)

func NewJenkinsWithTestData() *Jenkins {
	var auth Auth
	return NewJenkins(&auth, "http://example.com")
}

func Test(t *testing.T) {
	jenkins := NewJenkinsWithTestData()

	jobs, err := jenkins.GetJobs()

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(jobs) == 0 {
		t.Errorf("return no jobs\n")
	}
}
