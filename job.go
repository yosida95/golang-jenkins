package gojenkins

import "encoding/xml"

type Artifact struct {
	DisplayPath  string `json:"displayPath"`
	FileName     string `json:"fileName"`
	RelativePath string `json:"relativePath"`
}

type ScmAuthor struct {
	FullName    string `json:"fullName"`
	AbsoluteUrl string `json:"absoluteUrl"`
}

type ScmChangeSetPath struct {
	EditType string `json:"editType"`
	File     string `json:"File"`
}

type ChangeSetItem struct {
	AffectedPaths []string           `json:"affectedPaths"`
	CommitId      string             `json:"commitId"`
	Timestamp     int                `json:"timestamp"`
	Author        ScmAuthor          `json:"author"`
	Comment       string             `json:"comment"`
	Date          string             `json:"date"`
	Id            string             `json:"id"`
	Message       string             `json:"msg"`
	Paths         []ScmChangeSetPath `json:"paths"`
}

type ScmChangeSet struct {
	Kind  string          `json:"kind"`
	Items []ChangeSetItem `json:"items"`
}

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

	Artifacts []Artifact `json:"artifacts"`
	Actions   []Action   `json:"actions"`

	ChangeSet ScmChangeSet `json:"changeSet"`
}

type UpstreamCause struct {
	ShortDescription string `json:"shortDescription"`
	UpstreamBuild    int    `json:"upstreamBuild"`
	UpstreamProject  string `json:"upstreamProject"`
	UpstreamUrl      string `json:"upstreamUrl"`
}

type Job struct {
	Class   string   `json:"_class"`
	Actions []Action `json:"actions"`
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Color   string   `json:"color"`

	Buildable    bool     `json:"buildable"`
	Builds       []Build  `json:"builds"`
	DisplayName  string   `json:"displayName"`
	Description  string   `json:"description"`
	HealthReport []Health `json:"healthReport"`

	LastCompletedBuild    Build `json:"lastCompletedBuild"`
	LastFailedBuild       Build `json:"lastFailedBuild"`
	LastStableBuild       Build `json:"lastStableBuild"`
	LastSuccessfulBuild   Build `json:"lastSuccessfulBuild"`
	LastUnstableBuild     Build `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild Build `json:"lastUnsuccessfulBuild"`

	Property []Property `json:"property"`
}

type Health struct {
	Description string `json:"description"`
}

type Property struct {
	Parameters []JobParameter `json:"parameterDefinitions"`
}

type JobParameter struct {
	Default     Parameter `json:"defaultParameterValue"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Choices     []string  `json:"choices"`
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
	MavenName                        string               `xml:"mavenName"`
	MavenValidationLevel             string               `xml:"mavenValidationLevel"`
	DefaultGoals                     string               `xml:"defaultGoals"`
	RunHeadless                      string               `xml:"runHeadless"`
	DisableTriggerDownstreamProjects string               `xml:"disableTriggerDownstreamProjects"`
	Settings                         JobSettings          `xml:"settings"`
	GlobalSettings                   JobSettings          `xml:"globalSettings"`
	RunPostStepsIfResult             RunPostStepsIfResult `xml:"runPostStepsIfResult"`
	Postbuilders                     PostBuilders         `xml:"postbuilders"`
}

type Scm struct {
	ScmContent
	Class  string `xml:"class,attr"`
	Plugin string `xml:"plugin,attr"`
}

type ScmContent interface{}

type ScmSvn struct {
	Locations              Locations        `xml:"locations"`
	ExcludedRegions        string           `xml:"excludedRegions"`
	IncludedRegions        string           `xml:"includedRegions"`
	ExcludedUsers          string           `xml:"excludedUsers"`
	ExcludedRevprop        string           `xml:"excludedRevprop"`
	ExcludedCommitMessages string           `xml:"excludedCommitMessages"`
	WorkspaceUpdater       WorkspaceUpdater `xml:"workspaceUpdater"`
	IgnoreDirPropChanges   string           `xml:"ignoreDirPropChanges"`
	FilterChangelog        string           `xml:"filterChangelog"`
}

type WorkspaceUpdater struct {
	Class string `xml:"class,attr"`
}
type Locations struct {
	Location []ScmSvnLocation `xml:"hudson.scm.SubversionSCM_-ModuleLocation"`
}

type ScmSvnLocation struct {
	Remote                string `xml:"remote"`
	Local                 string `xml:"local"`
	DepthOption           string `xml:"depthOption"`
	IgnoreExternalsOption string `xml:"ignoreExternalsOption"`
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

type ScmGit struct {
	UserRemoteConfigs                 UserRemoteConfigs `xml:"userRemoteConfigs"`
	Branches                          Branches          `xml:"branches"`
	DoGenerateSubmoduleConfigurations bool              `xml:"doGenerateSubmoduleConfigurations"`
	GitBrowser                        GitBrowser        `xml:"browser"`
	GitSubmoduleCfg                   GitSubmoduleCfg   `xml:"submoduleCfg"`
	GitExtensions                     GitExtensions     `xml:"extensions"`
}

type UserRemoteConfigs struct {
	UserRemoteConfig UserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}

type UserRemoteConfig struct {
	Urls []string `xml:"url"`
}

type Branches struct {
	BranchesSpec []BranchesSpec `xml:"hudson.plugins.git.BranchSpec"`
}

type BranchesSpec struct {
	Name string `xml:"name"`
}

type GitBrowser struct {
	Class       string `xml:"class,attr"`
	Url         string `xml:"url"`
	ProjectName string `xml:"projectName"`
}

type GitSubmoduleCfg struct {
	Class string `xml:"class,attr"`
}

type GitExtensions struct {
	Class       string      `xml:"class,attr"`
	LocalBranch LocalBranch `xml:"hudson.plugins.git.extensions.impl.LocalBranch"`
}

type LocalBranch struct {
	LocalBranch string `xml:"localBranch"`
}

//UnmarshalXML implements xml.UnmarshalXML intrface
//Decode between multiple types of Scm. for now only SVN is supported
func (iscm *Scm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, v := range start.Attr {
		if v.Name.Local == "class" {
			iscm.Class = v.Value
		} else if v.Name.Local == "plugin" {
			iscm.Plugin = v.Value
		}
	}
	switch iscm.Class {
	case "hudson.scm.SubversionSCM":
		iscm.ScmContent = &ScmSvn{}
		err := d.DecodeElement(&iscm.ScmContent, &start)
		if err != nil {
			return err
		}
	case "hudson.plugins.git.GitSCM":
		iscm.ScmContent = &ScmGit{}
		err := d.DecodeElement(&iscm.ScmContent, &start)
		if err != nil {
			return err
		}
	}
	return nil
}
