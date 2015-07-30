package gojenkins

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&DockerSuite{})

func (s *DockerSuite) TestGettingJobs(c *C) {
	jenkins := NewJenkinsWithTestData()

	jobs, err := jenkins.GetJobs()

	c.Assert(err, IsNil)
	c.Assert(len(jobs), Not(Equals), 0)
}
