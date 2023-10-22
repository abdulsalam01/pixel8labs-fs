package github

type BigCacheResource interface {
	Set(key string, entry []byte) error
	Get(key string) ([]byte, error)
}

type UserRepo struct {
	bigCache BigCacheResource
}
