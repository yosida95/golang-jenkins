package gojenkins

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
