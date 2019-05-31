package gojenkins

import (
	"fmt"
	"net/url"
	"strings"
)

//InstallPlugin installs a given plugin to jenkins instance
func (jenkins *Jenkins) InstallPlugin(name, version string) error {

	payload := `<jenkins><install plugin="` + name + `@` + version + `" /></jenkins>`

	xmlBody := strings.NewReader(payload)

	var body interface{}
	err := jenkins.postXml("/pluginManager/installNecessaryPlugins", url.Values{}, xmlBody, &body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	return nil
}
