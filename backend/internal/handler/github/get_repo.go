package github

import (
	"net/http"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (h *Handler) GetRepo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		resp []entity.Repository
		err  error

		ctx = r.Context()
	)

	token := r.Header.Get(constant.HTTPHeaderAuthorization)
	resp, err = h.usecaseResource.GetRepo(ctx, entity.BaseRequestParams{
		Token: token,
	})
	return resp, err
}
