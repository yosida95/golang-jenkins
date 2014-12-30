package gojenkins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Auth struct {
	Username string
	ApiToken string
}

type Jenkins struct {
	auth    *Auth
	baseUrl string
}

func NewJenkins(auth *Auth, baseUrl string) *Jenkins {
	return &Jenkins{
		auth:    auth,
		baseUrl: baseUrl,
	}
}

func (jenkins *Jenkins) buildUrl(path string, params url.Values) (requestUrl string) {
	requestUrl = jenkins.baseUrl + path + "/api/json"
	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			requestUrl = requestUrl + "?" + queryString
		}
	}

	return
}

func (jenkins *Jenkins) sendRequest(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(jenkins.auth.Username, jenkins.auth.ApiToken)
	return http.DefaultClient.Do(req)
}

func (jenkins *Jenkins) parseResponse(resp *http.Response, body interface{}) (err error) {
	defer resp.Body.Close()

	if body == nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(data, body)
}

func (jenkins *Jenkins) get(path string, params url.Values, body interface{}) (err error) {
	requestUrl := jenkins.buildUrl(path, params)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	resp, err := jenkins.sendRequest(req)
	if err != nil {
		return
	}

	return jenkins.parseResponse(resp, body)
}

func (jenkins *Jenkins) post(path string, params url.Values, body interface{}) (err error) {
	requestUrl := jenkins.buildUrl(path, params)
	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return
	}

	resp, err := jenkins.sendRequest(req)
	if err != nil {
		return
	}

	return jenkins.parseResponse(resp, body)
}

// GetJobs returns all jobs you can read.
func (jenkins *Jenkins) GetJobs() (jobs []Job, err error) {
	var payload = struct {
		Jobs []Job `json:"jobs"`
	}{
		Jobs: jobs,
	}
	err = jenkins.get("", nil, &payload)
	return
}

// GetJob returns a job which has specified name.
func (jenkins *Jenkins) GetJob(name string) (job Job, err error) {
	err = jenkins.get(fmt.Sprintf("/job/%s", name), nil, &job)
	return
}

// GetBuild returns a number-th build result of specified job.
func (jenkins *Jenkins) GetBuild(job Job, number int) (build Build, err error) {
	err = jenkins.get(fmt.Sprintf("/job/%s/%d", job.Name, number), nil, &build)
	return
}

// Create a new build for this job.
// Params can be nil.
func (jenkins *Jenkins) Build(job Job, params url.Values) error {
	if params == nil {
		return jenkins.post(fmt.Sprintf("/job/%s/build", job.Name), params, nil)
	} else {
		return jenkins.post(fmt.Sprintf("/job/%s/buildWithParameters", job.Name), params, nil)
	}
}
