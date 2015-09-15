package gojenkins

import (
	"fmt"
	"testing"
)

func TestGetArtifacts(t *testing.T) {
	jenkins := NewJenkinsWithTestData()
	job, err := jenkins.GetJob("job_name_example")
	fmt.Println("job is %+v", job)
	if err != nil {
		fmt.Println(err)
		t.Errorf("job is not fetched")
	}
	var artifacts Artifacts
	err = job.getArtifacts(jenkins, &artifacts)
	if err != nil {
		t.Errorf("Falied to get Artifacts")
	}
	t.Log("Job artifacts are fetched successfully")
}
