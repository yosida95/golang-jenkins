package gojenkins

import (
	"net/url"
	"strings"
)

//InstallPlugin installs a given plugin to jenkins instance
func (jenkins *Jenkins) InstallPlugin(name, version string) error {

	payload := `<jenkins><install plugin="` + name + `@` + version + `" /></jenkins>`

	xmlBody := strings.NewReader(payload)

	var body interface{}
	return jenkins.postXml("/pluginManager/installNecessaryPlugins", url.Values{}, xmlBody, &body)
}
