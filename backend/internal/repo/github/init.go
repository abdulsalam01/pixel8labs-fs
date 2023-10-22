package github

func New(bigCache BigCacheResource) *UserRepo {
	return &UserRepo{
		bigCache: bigCache,
	}
}
