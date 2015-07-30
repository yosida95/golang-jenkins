package gojenkins

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/samalba/dockerclient"
	. "gopkg.in/check.v1"
)

const (
	JENKINS_CONTAINER_NAME  = "jenkins-test"
	JENKINS_TEST_IMAGE_NAME = "jenkins-test-image"
)

type DockerSuite struct{}

var docker *dockerclient.DockerClient

func (s *DockerSuite) SetUpSuite(c *C) {
	var tlsc tls.Config
	cert, _ := tls.LoadX509KeyPair(os.Getenv("DOCKER_CERT_PATH")+"/cert.pem", os.Getenv("DOCKER_CERT_PATH")+"/key.pem")
	tlsc.Certificates = append(tlsc.Certificates, cert)
	tlsc.InsecureSkipVerify = true
	docker, _ = dockerclient.NewDockerClient(os.Getenv("DOCKER_HOST"), &tlsc)

	buildTestJenkinsImage()
	containerId, err := createJenkinsContainer()
	if err != nil {
		fmt.Printf("error creating container: %s", err)
		c.FailNow()
	}

	if err := docker.StartContainer(containerId, &dockerclient.HostConfig{PublishAllPorts: true}); err != nil {
		fmt.Printf("error starting container: %s", err)
		c.FailNow()
	}
	c.Assert(waitForJenkinsToStartup(), Equals, true)
}

func (s *DockerSuite) TearDownSuite(c *C) {
	docker.KillContainer(JENKINS_CONTAINER_NAME, "9")
	docker.RemoveContainer(JENKINS_CONTAINER_NAME, false, true)
}

func NewJenkinsWithTestData() *Jenkins {
	containerInfo, _ := docker.InspectContainer(JENKINS_CONTAINER_NAME)
	portMapping := containerInfo.NetworkSettings.Ports["8080/tcp"][0].HostPort
	hostIp := findIp(os.Getenv("DOCKER_HOST"))
	var auth Auth

	return NewJenkins(&auth, fmt.Sprintf("http://%s:%s", hostIp, portMapping))
}

func waitForJenkinsToStartup() bool {
	logMessageIndicatingJenkinsHasStarted := "Jenkins is fully up and running"
	reader, _ := docker.ContainerLogs(JENKINS_CONTAINER_NAME, &dockerclient.LogOptions{Stdout: true, Stderr: true, Follow: true})
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	success := make(chan bool)

	go func() {
		for scanner.Scan() == true {
			if strings.Contains(scanner.Text(), logMessageIndicatingJenkinsHasStarted) {
				success <- true
				break
			}
		}
	}()

	go func() {
		for i := 0; i < 6; i++ {
			time.Sleep(5 * time.Second)
		}
		success <- false
	}()

	return <-success
}

func createJenkinsContainer() (string, error) {
	containerConfig := &dockerclient.ContainerConfig{
		Image:       JENKINS_TEST_IMAGE_NAME + ":latest",
		AttachStdin: false,
		Tty:         false}
	return docker.CreateContainer(containerConfig, JENKINS_CONTAINER_NAME)
}

func findIp(input string) string {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}

func buildTestJenkinsImage() (string, error) {
	fmt.Println("Building jenkins image with test data")

	dockerFile, err := os.Open("./docker/Dockerfile.tar")
	defer dockerFile.Close()

	reader, err := docker.BuildImage(&dockerclient.BuildImage{
		Context:        dockerFile,
		RepoName:       JENKINS_TEST_IMAGE_NAME,
		SuppressOutput: false,
		Remove:         true,
	})
	defer reader.Close()

	if err != nil {
		fmt.Printf("error building image: %s", err)
		return "", err
	}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() == true {
		fmt.Println(scanner.Text())
	}

	return JENKINS_TEST_IMAGE_NAME, nil
}
