package github

import (
	"context"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (u *Usecase) GetRepo(ctx context.Context, params entity.BaseRequestParams) ([]entity.Repository, error) {
	var (
		resp []entity.Repository
		err  error
	)

	// Get repo list.
	resp, err = u.githubResource.GetRepoByUser(ctx, entity.BaseRequestParams{
		Token:    params.Token,
		Username: constant.DefaultUser,
	})
	if err != nil {
		return resp, err
	}
	if len(resp) > 6 {
		resp = resp[:6] // Get only max 6.
	}

	return resp, nil
}
