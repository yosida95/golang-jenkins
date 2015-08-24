package gojenkins

import "encoding/xml"

type ListView struct {
	XMLName         xml.Name `xml:"hudson.model.ListView"`
	Name            string   `xml:"name"`
	FilterExecutors bool     `xml:"filterExecutors"`
	FilterQueue     bool     `xml:"filterQueue"`
	Columns         Columns  `xml:"columns"`
}

func NewListView(name string) ListView {
	columns := Columns{Column: []Column{StatusColumn{}, WeatherColumn{}, JobColumn{}, LastSuccessColumn{}, LastFailureColumn{}, LastDurationColumn{}, BuildButtonColumn{}}}
	return ListView{Name: name, FilterExecutors: false, FilterQueue: false, Columns: columns}
}

type Column interface {
}

type Columns struct {
	XMLName xml.Name `xml:"columns"`
	Column  []Column
}

type StatusColumn struct {
	XMLName xml.Name `xml:"hudson.view.StatusColumn"`
}
type WeatherColumn struct {
	XMLName xml.Name `xml:"hudson.view.WeatherColumn"`
}

type JobColumn struct {
	XMLName xml.Name `xml:"hudson.view.JobColumn"`
}
type LastSuccessColumn struct {
	XMLName xml.Name `xml:"hudson.view.LastSuccessColumn"`
}
type LastFailureColumn struct {
	XMLName xml.Name `xml:"hudson.view.LastFailureColumn"`
}
type LastDurationColumn struct {
	XMLName xml.Name `xml:"hudson.view.LastDurationColumn"`
}
type BuildButtonColumn struct {
	XMLName xml.Name `xml:"hudson.view.BuildButtonColumn"`
}
