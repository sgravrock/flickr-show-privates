package auth

import (
	"net/http"

	"github.com/mrjones/oauth"
)

type UiAdapter interface {
	PromptForAccessCode(url string) (string, error)
}

type Authenticator interface {
	Authenticate() (*http.Client, error)
}

func NewAuthenticator(key string, secret string,
	oauthClient OauthClient, ui UiAdapter) Authenticator {

	if oauthClient == nil {
		oauthClient = NewOauthClient()
	}
	if ui == nil {
		ui = NewConsoleUiAdapter()
	}

	return &defaultAuthenticator{key, secret, oauthClient, ui}
}

type defaultAuthenticator struct {
	key         string
	secret      string
	oauthClient OauthClient
	ui          UiAdapter
}

func (a *defaultAuthenticator) Authenticate() (*http.Client, error) {

	consumer := a.oauthClient.NewConsumer(
		a.key,
		a.secret,
		"https://www.flickr.com/services/oauth/request_token",
		"https://www.flickr.com/services/oauth/authorize",
		"https://www.flickr.com/services/oauth/access_token")
	consumer.SetAdditionalParams(map[string]string{
		"perms": "read",
	})

	accessToken, err := a.getAccessToken(consumer)
	if err != nil {
		return nil, err
	}

	return consumer.MakeHttpClient(accessToken)
}

func (a *defaultAuthenticator) getAccessToken(consumer OauthConsumer) (
	*oauth.AccessToken, error) {

	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return nil, err
	}

	accessCode, err := a.ui.PromptForAccessCode(url)
	if err != nil {
		return nil, err
	}

	return consumer.AuthorizeToken(requestToken, accessCode)
}
