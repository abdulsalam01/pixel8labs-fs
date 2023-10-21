package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abdulsalam/pixel8labs/helper"
	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
	"golang.org/x/oauth2"
)

func (s *Service) GetLogin(ctx context.Context) string {
	// Redirect user to GitHub for OAuth login.
	authURL := s.OAuthConfig.AuthCodeURL(constant.DefaultString, oauth2.AccessTypeOffline)
	return authURL
}

func (s *Service) GetLogout(ctx context.Context, params entity.BaseRequestParams) (bool, error) {
	var (
		result bool
		err    error
	)

	// Create an HTTP client.
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s/%s/%s", s.Config.BaseUrl, "applications", s.Config.ClientID, "token")

	// Initialize the params.
	bodyParams := map[string]string{"access_token": params.Token}
	bodyByte, err := json.Marshal(bodyParams)
	if err != nil {
		return result, err
	}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		log.Fatal("Error creating request:", err)
		return result, err
	}

	// Set a custom header (e.g., an authorization header).
	authHeader := helper.Base64Encode(s.Config.ClientID + ":" + s.Config.ClientSecret)
	req.Header.Set(constant.HTTPHeaderAuthorization, "Basic "+authHeader)
	req.Header.Set(constant.HTTPHeaderAPIVersion, constant.DefaultAPIVersionValue)
	req.Header.Set(constant.HTTPHeaderAcceptContent, constant.DefaultAPIAcceptContentValue)

	// Send the request.
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return result, nil
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return result, nil
	}

	return true, nil
}

func (s *Service) GetCallback(ctx context.Context, code string) (*oauth2.Token, error) {
	// Handle OAuth callback.
	token, err := s.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		return token, err
	}

	return token, nil
}
