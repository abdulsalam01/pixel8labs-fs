package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (s *Service) GetRepoByUser(ctx context.Context, params entity.BaseRequestParams) ([]entity.Repository, error) {
	var (
		repos []entity.Repository
		err   error
	)

	// Create an HTTP client.
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", s.Config.BaseUrl, "user/repos")

	// Set to default user.
	if params.Token == constant.DefaultString {
		url = fmt.Sprintf("%s/%s/%s/repos", s.Config.BaseUrl, "users", params.Username)
	}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return repos, err
	}

	// Set to logged in user.
	if params.Token != constant.DefaultString {
		req.Header.Set(constant.HTTPHeaderAuthorization, params.Token)
	}

	// Set a custom header (e.g., an authorization header).
	req.Header.Set(constant.HTTPHeaderAPIVersion, constant.DefaultAPIVersionValue)
	req.Header.Set(constant.HTTPHeaderAcceptContent, constant.DefaultAPIAcceptContentValue)

	// Send the request.
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return repos, nil
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return repos, nil
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return repos, nil
	}

	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return repos, err
	}

	return repos, nil
}
