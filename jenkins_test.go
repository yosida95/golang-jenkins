package gojenkins

import (
	"fmt"
	"testing"
	"time"
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

func TestCreateJobItem(t *testing.T) {
	jenkins := NewJenkinsWithTestData()
	scm := Scm{
		Locations: Locations{
			[]Location{
				ScmSvnLocation{IgnoreExternalsOption: "false", DepthOption: "infinity", Local: ".", Remote: "http://some-svn-url"},
			},
		},
		Class:                "hudson.scm.SubversionSCM",
		Plugin:               "subversion@1.54",
		IgnoreDirPropChanges: "false",
		FilterChanglog:       "false",
		WorkspaceUpdater:     WorkspaceUpdater{Class: "hudson.scm.subversion.UpdateUpdater"},
	}
	triggers := Triggers{[]Trigger{ScmTrigger{}}}
	postStep := RunPostStepsIfResult{Name: "FAILURE", Ordinal: "2", Color: "RED", CompleteBuild: "true"}
	settings := JobSettings{Class: "jenkins.mvn.DefaultSettingsProvider"}
	globalSettings := JobSettings{Class: "jenkins.mvn.DefaultSettingsProvider"}
	jobItem := MavenJobItem{
		Plugin:               "maven-plugin@2.7.1",
		Description:          "test description",
		Scm:                  scm,
		Triggers:             triggers,
		RunPostStepsIfResult: postStep,
		Settings:             settings,
		GlobalSettings:       globalSettings,
	}

	newJobName := fmt.Sprintf("test-%d", time.Now().UnixNano())
	err := jenkins.CreateJob(jobItem, newJobName)

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	jobs, _ := jenkins.GetJobs()
	foundNewJob := false
	for _, v := range jobs {
		if v.Name == newJobName {
			foundNewJob = true
		}
	}

	if !foundNewJob {
		t.Errorf("error %s not found\n", newJobName)
	}
}
