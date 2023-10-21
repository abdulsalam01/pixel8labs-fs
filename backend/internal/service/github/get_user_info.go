package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/abdulsalam/pixel8labs/internal/constant"
	"github.com/abdulsalam/pixel8labs/internal/entity"
)

func (s *Service) GetInfoByUsername(ctx context.Context, params entity.BaseRequestParams) (entity.Account, error) {
	var (
		account entity.Account
		email   []entity.Email
		err     error

		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		account, err = s.GetAccount(ctx, params)
	}()

	wg.Wait()
	if err != nil {
		return account, err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		email, err = s.GetEmail(ctx, params)
	}()

	wg.Wait()
	if err != nil {
		return account, err
	}

	// Get email.
	if len(email) > 0 {
		account.Email = email[0].Email
	}

	return account, nil
}

func (s *Service) GetAccount(ctx context.Context, params entity.BaseRequestParams) (entity.Account, error) {
	var (
		account entity.Account
		err     error
	)

	// Create an HTTP client.
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", s.Config.BaseUrl, "user")

	// Set to default user.
	if params.Token == constant.DefaultString {
		url = fmt.Sprintf("%s/%s/%s", s.Config.BaseUrl, "users", params.Username)
	}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return account, err
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
		return account, nil
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return account, nil
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return account, nil
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return account, err
	}

	return account, nil
}
func (s *Service) GetEmail(ctx context.Context, params entity.BaseRequestParams) ([]entity.Email, error) {
	var (
		account = []entity.Email{
			{Email: constant.DefaultMail},
		}
		err error
	)

	// Create an HTTP client.
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", s.Config.BaseUrl, "user/emails")

	// Set to default user.
	if params.Token == constant.DefaultString {
		return account, nil
	}

	// Create a request with a custom header.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return account, err
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
		return account, nil
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return account, nil
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return account, nil
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return account, err
	}

	return account, nil
}
