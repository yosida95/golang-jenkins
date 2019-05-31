package gojenkins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SecretCreds struct {
	Scope       string `json:"scope"`
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	Description string `json:"description"`
	Class       string `json:"$class"`
}

type UsernamePasswordCredential struct {
	Scope       string `json:"scope"`
	User        string `json:"username"`
	Password    string `json:"password"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Class       string `json:"$class"`
}

func (jenkins *Jenkins) CreateCredentialsSecret(scope, id, description, secret string) {
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	m := SecretCreds{
		Scope:       scope,
		ID:          id,
		Secret:      secret,
		Description: description,
		Class:       "org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	jsonPayload := `{"": "0", "credentials":` + b.String() + `}`
	jenkins.createCredentials(header, jsonPayload)
}

func (jenkins *Jenkins) postForm(path string, additionalHeader map[string]string, params url.Values, data url.Values, body interface{}) (err error) {
	requestURL := jenkins.baseUrl + path
	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			requestURL = requestURL + "?" + queryString
		}
	}

	req, err := http.NewRequest("POST", requestURL, strings.NewReader(data.Encode()))

	for k, v := range additionalHeader {
		req.Header.Add(k, v)
	}
	if _, err := jenkins.checkCrumb(req); err != nil {
		return err
	}

	resp, err := jenkins.sendRequest(req)
	if err != nil {
		return
	}
	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return errors.New(fmt.Sprintf("error: HTTP POST returned status code %d (expected 2xx)", resp.StatusCode))
	}

	return jenkins.parseResponse(resp, body)
}

func (jenkins *Jenkins) CreateUsernamePasswordCredential(scope, id, user, description, password string) {
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	m := UsernamePasswordCredential{
		Scope:       scope,
		ID:          id,
		User:        user,
		Password:    password,
		Description: description,
		Class:       "com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	jsonPayload := `{"": "0", "credentials":` + b.String() + `}`
	jenkins.createCredentials(header, jsonPayload)
}

func (jenkins *Jenkins) createCredentials(header map[string]string, jsonPayload string) {
	form := url.Values{"json": []string{jsonPayload}}

	var body interface{}
	err := jenkins.postForm("/credentials/store/system/domain/_/createCredentials", header, url.Values{}, form, body)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println(body)
	}
}
