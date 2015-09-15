package gojenkins

import (
	"fmt"
	"testing"
)

func TestGetArtifacts(t *testing.T) {
	jenkins := NewJenkinsWithTestData()
	job, err := jenkins.GetJob("live-neo-unit-tests")
	fmt.Println("job is %+v", job)
	if err != nil {
		fmt.Println(err)
	}
	var artifacts Artifacts
	job.getArtifacts(jenkins, &artifacts)
	fmt.Println("job Artifacts are  %+v", artifacts)
}
