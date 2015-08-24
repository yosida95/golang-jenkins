package gojenkins

import "encoding/xml"

type Build struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	Url    string `json:"url"`

	FullDisplayName string `json:"fullDisplayName"`
	Description     string `json:"description"`

	Timestamp         int `json:"timestamp"`
	Duration          int `json:"duration"`
	EstimatedDuration int `json:"estimatedDuration"`

	Building bool   `json:"building"`
	KeepLog  bool   `json:"keepLog"`
	Result   string `json:"result"`
}

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`

	Buildable   bool   `json:"buildable"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`

	LastCompletedBuild    Build `json:"lastCompletedBuild"`
	LastFailedBuild       Build `json:"lastFailedBuild"`
	LastStableBuild       Build `json:"lastStableBuild"`
	LastSuccessfulBuild   Build `json:"lastSuccessfulBuild"`
	LastUnstableBuild     Build `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild Build `json:"lastUnsuccessfulBuild"`
}

type MavenJobItem struct {
	XMLName                          struct{}             `xml:"maven2-moduleset"`
	Plugin                           string               `xml:"plugin,attr"`
	Actions                          string               `xml:"actions"`
	Description                      string               `xml:"description"`
	KeepDependencies                 string               `xml:"keepDependencies"`
	Properties                       JobProperties        `xml:"properties"`
	Scm                              Scm                  `xml:"scm"`
	CanRoam                          string               `xml:"canRoam"`
	Disabled                         string               `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding string               `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   string               `xml:"blockBuildWhenUpstreamBuilding"`
	Triggers                         Triggers             `xml:"triggers"`
	ConcurrentBuild                  string               `xml:"concurrentBuild"`
	Goals                            string               `xml:"goals"`
	AggregatorStyleBuild             string               `xml:"aggregatorStyleBuild"`
	IncrementalBuild                 string               `xml:"incrementalBuild"`
	IgnoreUpstremChanges             string               `xml:"ignoreUpstremChanges"`
	ArchivingDisabled                string               `xml:"archivingDisabled"`
	SiteArchivingDisabled            string               `xml:"siteArchivingDisabled"`
	FingerprintingDisabled           string               `xml:"fingerprintingDisabled"`
	ResolveDependencies              string               `xml:"resolveDependencies"`
	ProcessPlugins                   string               `xml:"processPlugins"`
	MavenValidationLevel             string               `xml:"mavenValidationLevel"`
	RunHeadless                      string               `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string               `xml:"disableTriggerDownstreamProjects"`
	Settings                         JobSettings          `xml:"settings"`
	GlobalSettings                   JobSettings          `xml:"globalSettings"`
	RunPostStepsIfResult             RunPostStepsIfResult `xml:"runPostStepsIfResult"`
	Postbuilders                     PostBuilders         `xml:"postbuilders"`
}

type Scm struct {
	Locations              Locations        `xml:"locations"`
	Class                  string           `xml:"class,attr"`
	Plugin                 string           `xml:"plugin,attr"`
	ExcludedRegions        string           `xml:"excludedRegions"`
	IncludedRegions        string           `xml:"includedRegions"`
	ExcludedUsers          string           `xml:"excludedUsers"`
	ExcludedRevprop        string           `xml:"excludedRevprop"`
	ExcludedCommitMessages string           `xml:"excludedCommitMessages"`
	WorkspaceUpdater       WorkspaceUpdater `xml:"workspaceUpdater"`
	IgnoreDirPropChanges   string           `xml:"ignoreDirPropChanges"`
	FilterChanglog         string           `xml:"filterChangelog"`
}

type WorkspaceUpdater struct {
	Class string `xml:"class,attr"`
}
type Locations struct {
	Location []Location
}
type Location interface {
}
type ScmSvnLocation struct {
	XMLName               struct{} `xml:"hudson.scm.SubversionSCM_-ModuleLocation"`
	Remote                string   `xml:"remote"`
	Local                 string   `xml:"local"`
	DepthOption           string   `xml:"depthOption"`
	IgnoreExternalsOption string   `xml:"ignoreExternalsOption"`
}

type PostBuilders struct {
	XMLName     xml.Name `xml:"postbuilders"`
	PostBuilder []PostBuilder
}

type PostBuilder interface {
}

type ShellBuilder struct {
	XMLName xml.Name `xml:"hudson.tasks.Shell"`
	Command string   `xml:"command"`
}

type JobSettings struct {
	Class      string `xml:"class,attr"`
	JobSetting []JobSetting
}

type JobSetting struct {
}
type JobProperties struct {
}
type Triggers struct {
	Trigger []Trigger
}
type Trigger interface {
}
type ScmTrigger struct {
	XMLName               struct{} `xml:"hudson.triggers.SCMTrigger"`
	Spec                  string   `xml:"spec"`
	IgnorePostCommitHooks string   `xml:"ignorePostCommitHooks"`
}
type RunPostStepsIfResult struct {
	Name          string `xml:"name"`
	Ordinal       string `xml:"ordinal"`
	Color         string `xml:"color"`
	CompleteBuild string `xml:"completeBuild"`
}
