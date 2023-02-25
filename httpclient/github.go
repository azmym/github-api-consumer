package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github-api-consumer/config"
	"github-api-consumer/domain"
	"io"
	"net/http"
)

const (
	githubHost         = "https://api.github.com"
	acceptKey          = "Accept"
	acceptValue        = "application/vnd.github+json"
	authorizationKey   = "Authorization"
	authorizationValue = "Bearer %s"

	languagesEndpoint    = "%s/repos/%s/%s/languages"
	repositoriesEndpoint = "%s/user/repos"
	userEndpoint         = "%s/user"
)

type (
	GithubClient struct {
		*http.Client
	}
)

func (g *GithubClient) GetListAllUserRepos() (*[]domain.Repo, error) {
	//prepare the request
	request, err := prepareRequest(http.MethodGet, githubHost, repositoriesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	//fire the request
	response, err := g.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	repos := make([]domain.Repo, 0)
	err = json.Unmarshal(responseBody, &repos)
	if err != nil {
		return nil, err
	}
	return &repos, nil
}

func (g *GithubClient) GetUserInformation() (*domain.User, error) {
	//prepare the request
	request, err := prepareRequest(http.MethodGet, githubHost, userEndpoint, nil)
	if err != nil {
		return nil, err
	}
	//fire the request
	response, err := g.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	user := new(domain.User)
	err = json.Unmarshal(responseBody, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (g *GithubClient) GetRepoLanguages(user string, repoName string) (*domain.Languages, error) {
	endpoint := fmt.Sprintf(languagesEndpoint, "%s", user, repoName)
	request, err := prepareRequest(http.MethodGet, githubHost, endpoint, nil)
	if err != nil {
		return nil, err
	}
	//fire the request
	response, err := g.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	languages := make(domain.Languages)
	err = json.Unmarshal(responseBody, &languages)
	if err != nil {
		return nil, err
	}
	return &languages, nil
}

// prepare the request
func prepareRequest(method, host, endpoint string, body []byte) (*http.Request, error) {
	//build a header
	header := buildHeader()
	//build a request
	request, err := buildRequest(method, host, endpoint, nil)
	if err != nil {
		return nil, err
	}
	//assign header to the request
	request.Header = header
	return request, nil
}

// build header for the http request
func buildHeader() http.Header {
	header := make(http.Header)
	header.Set(acceptKey, acceptValue)
	auth := fmt.Sprintf(authorizationValue, config.EnvVariable.GithubToken)
	header.Set(authorizationKey, auth)
	return header
}

// build request based on method type
func buildRequest(method, host, endpoint string, body []byte) (*http.Request, error) {
	url := fmt.Sprintf(endpoint, host)
	if http.MethodGet == method || nil == body {
		return http.NewRequest(method, url, http.NoBody)
	}
	return http.NewRequest(method, url, bytes.NewReader(body))
}
