package github

import (
	"net/http"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		resp entity.Account
		err  error

		ctx = r.Context()
	)

	token := r.Header.Get(constant.HTTPHeaderAuthorization)
	resp, err = h.usecaseResource.GetUser(ctx, entity.BaseRequestParams{
		Token: token,
	})
	return resp, err
}
