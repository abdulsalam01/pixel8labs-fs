package github

import (
	"context"

	"github.com/abdulsalam/pixel8labs/internal/entity"
	"golang.org/x/oauth2"
)

type GithubResource interface {
	GetInfoByUsername(ctx context.Context, params entity.BaseRequestParams) (entity.Account, error)
	GetRepoByUser(ctx context.Context, params entity.BaseRequestParams) ([]entity.Repository, error)

	// Login.
	GetLogin(ctx context.Context) string
	GetLogout(ctx context.Context, params entity.BaseRequestParams) (bool, error)
	GetCallback(ctx context.Context, code string) (*oauth2.Token, error)
}

type UserResource interface {
	SetVisitorByID(ctx context.Context, id int64) (uint64, error)
}

type Usecase struct {
	githubResource GithubResource
	userResource   UserResource
}
