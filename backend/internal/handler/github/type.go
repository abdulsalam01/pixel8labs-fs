package github

import (
	"context"

	"github.com/abdulsalam/pixel8labs/internal/entity"
	"golang.org/x/oauth2"
)

type UsecaseResource interface {
	Auth(ctx context.Context, code string) (*oauth2.Token, error)
	LoginUrl(ctx context.Context) string
	LogoutUrl(ctx context.Context, params entity.BaseRequestParams) (bool, error)

	GetUser(ctx context.Context, params entity.BaseRequestParams) (entity.Account, error)
	GetRepo(ctx context.Context, params entity.BaseRequestParams) ([]entity.Repository, error)
}

type Handler struct {
	usecaseResource UsecaseResource
}
