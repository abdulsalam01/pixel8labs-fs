package github

func New(githubResource GithubResource, userResource UserResource) *Usecase {
	return &Usecase{
		githubResource: githubResource,
		userResource:   userResource,
	}
}
