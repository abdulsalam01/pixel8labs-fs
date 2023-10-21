package github

import (
	"context"
	"errors"
	"strings"

	"github.com/abdulsalam/pixel8labs/internal/entity"
	"golang.org/x/oauth2"
)

func (u *Usecase) LoginUrl(ctx context.Context) string {
	return u.githubResource.GetLogin(ctx)
}

func (u *Usecase) LogoutUrl(ctx context.Context, params entity.BaseRequestParams) (bool, error) {
	tokenize := strings.Split(params.Token, " ")
	if len(tokenize) < 2 {
		return false, errors.New("error with the token")
	}

	params.Token = tokenize[1]
	return u.githubResource.GetLogout(ctx, params)
}

func (u *Usecase) Auth(ctx context.Context, code string) (*oauth2.Token, error) {
	var (
		resp *oauth2.Token
		err  error
	)

	resp, err = u.githubResource.GetCallback(ctx, code)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
