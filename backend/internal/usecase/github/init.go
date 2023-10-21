package github

func New(githubResource GithubResource) *Usecase {
	return &Usecase{
		githubResource: githubResource,
	}
}
