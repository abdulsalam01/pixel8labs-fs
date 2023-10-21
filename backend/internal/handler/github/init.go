package github

func New(usecaseResource UsecaseResource) *Handler {
	return &Handler{
		usecaseResource: usecaseResource,
	}
}
