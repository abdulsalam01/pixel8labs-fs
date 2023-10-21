package github

import (
	"github.com/abdulsalam/pixel8labs/internal/entity"
	"golang.org/x/oauth2"
)

type Service struct {
	OAuthConfig oauth2.Config
	Config      entity.Config
}

func New(oauth oauth2.Config, config entity.Config) *Service {
	return &Service{
		Config:      config,
		OAuthConfig: oauth,
	}
}
