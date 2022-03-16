package uploader

type Uploader struct {
	GithubToken  string
	OwnerName    string
	RepoName     string
	SarifContent string
	PRNumber     int
}
