package github

import (
	"context"
	"net/http"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (h *Handler) GetCode(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		resp string
		ctx  = r.Context()
	)

	resp = h.usecaseResource.LoginUrl(ctx)
	ctx = context.WithValue(ctx, constant.CodeToken, resp) //nolint:all

	return resp, nil
}

func (h *Handler) GetLogout(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		resp bool
		err  error

		ctx = r.Context()
	)

	token := r.Header.Get(constant.HTTPHeaderAuthorization)
	resp, err = h.usecaseResource.LogoutUrl(ctx, entity.BaseRequestParams{
		Token: token,
	})
	return resp, err
}

// Callback for continue to direct to github oauth.
func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	code := r.URL.Query().Get(constant.CodeToken)

	token, err := h.usecaseResource.Auth(ctx, code)
	if err != nil {
		return token, err
	}

	ctx = context.WithValue(ctx, constant.AccessToken, token.AccessToken) //nolint:all
	return token, nil
}
