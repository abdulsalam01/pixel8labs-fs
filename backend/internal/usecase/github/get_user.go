package github

import (
	"context"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (u *Usecase) GetUser(ctx context.Context, params entity.BaseRequestParams) (entity.Account, error) {
	var (
		resp entity.Account
		err  error
	)

	// Get account detail.
	resp, err = u.githubResource.GetInfoByUsername(ctx, entity.BaseRequestParams{
		Token:    params.Token,
		Username: constant.DefaultUser,
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
