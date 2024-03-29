// This file was generated by counterfeiter
package authfakes

import (
	"net/http"
	"sync"

	"github.com/mrjones/oauth"
	"github.com/sgravrock/flickr-show-privates/auth"
)

type FakeOauthConsumer struct {
	GetRequestTokenAndUrlStub        func(callbackUrl string) (rtoken *oauth.RequestToken, loginUrl string, err error)
	getRequestTokenAndUrlMutex       sync.RWMutex
	getRequestTokenAndUrlArgsForCall []struct {
		callbackUrl string
	}
	getRequestTokenAndUrlReturns struct {
		result1 *oauth.RequestToken
		result2 string
		result3 error
	}
	AuthorizeTokenStub        func(rtoken *oauth.RequestToken, verificationCode string) (atoken *oauth.AccessToken, err error)
	authorizeTokenMutex       sync.RWMutex
	authorizeTokenArgsForCall []struct {
		rtoken           *oauth.RequestToken
		verificationCode string
	}
	authorizeTokenReturns struct {
		result1 *oauth.AccessToken
		result2 error
	}
	SetAdditionalParamsStub        func(params map[string]string)
	setAdditionalParamsMutex       sync.RWMutex
	setAdditionalParamsArgsForCall []struct {
		params map[string]string
	}
	MakeHttpClientStub        func(token *oauth.AccessToken) (*http.Client, error)
	makeHttpClientMutex       sync.RWMutex
	makeHttpClientArgsForCall []struct {
		token *oauth.AccessToken
	}
	makeHttpClientReturns struct {
		result1 *http.Client
		result2 error
	}
}

func (fake *FakeOauthConsumer) GetRequestTokenAndUrl(callbackUrl string) (rtoken *oauth.RequestToken, loginUrl string, err error) {
	fake.getRequestTokenAndUrlMutex.Lock()
	fake.getRequestTokenAndUrlArgsForCall = append(fake.getRequestTokenAndUrlArgsForCall, struct {
		callbackUrl string
	}{callbackUrl})
	fake.getRequestTokenAndUrlMutex.Unlock()
	if fake.GetRequestTokenAndUrlStub != nil {
		return fake.GetRequestTokenAndUrlStub(callbackUrl)
	} else {
		return fake.getRequestTokenAndUrlReturns.result1, fake.getRequestTokenAndUrlReturns.result2, fake.getRequestTokenAndUrlReturns.result3
	}
}

func (fake *FakeOauthConsumer) GetRequestTokenAndUrlCallCount() int {
	fake.getRequestTokenAndUrlMutex.RLock()
	defer fake.getRequestTokenAndUrlMutex.RUnlock()
	return len(fake.getRequestTokenAndUrlArgsForCall)
}

func (fake *FakeOauthConsumer) GetRequestTokenAndUrlArgsForCall(i int) string {
	fake.getRequestTokenAndUrlMutex.RLock()
	defer fake.getRequestTokenAndUrlMutex.RUnlock()
	return fake.getRequestTokenAndUrlArgsForCall[i].callbackUrl
}

func (fake *FakeOauthConsumer) GetRequestTokenAndUrlReturns(result1 *oauth.RequestToken, result2 string, result3 error) {
	fake.GetRequestTokenAndUrlStub = nil
	fake.getRequestTokenAndUrlReturns = struct {
		result1 *oauth.RequestToken
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOauthConsumer) AuthorizeToken(rtoken *oauth.RequestToken, verificationCode string) (atoken *oauth.AccessToken, err error) {
	fake.authorizeTokenMutex.Lock()
	fake.authorizeTokenArgsForCall = append(fake.authorizeTokenArgsForCall, struct {
		rtoken           *oauth.RequestToken
		verificationCode string
	}{rtoken, verificationCode})
	fake.authorizeTokenMutex.Unlock()
	if fake.AuthorizeTokenStub != nil {
		return fake.AuthorizeTokenStub(rtoken, verificationCode)
	} else {
		return fake.authorizeTokenReturns.result1, fake.authorizeTokenReturns.result2
	}
}

func (fake *FakeOauthConsumer) AuthorizeTokenCallCount() int {
	fake.authorizeTokenMutex.RLock()
	defer fake.authorizeTokenMutex.RUnlock()
	return len(fake.authorizeTokenArgsForCall)
}

func (fake *FakeOauthConsumer) AuthorizeTokenArgsForCall(i int) (*oauth.RequestToken, string) {
	fake.authorizeTokenMutex.RLock()
	defer fake.authorizeTokenMutex.RUnlock()
	return fake.authorizeTokenArgsForCall[i].rtoken, fake.authorizeTokenArgsForCall[i].verificationCode
}

func (fake *FakeOauthConsumer) AuthorizeTokenReturns(result1 *oauth.AccessToken, result2 error) {
	fake.AuthorizeTokenStub = nil
	fake.authorizeTokenReturns = struct {
		result1 *oauth.AccessToken
		result2 error
	}{result1, result2}
}

func (fake *FakeOauthConsumer) SetAdditionalParams(params map[string]string) {
	fake.setAdditionalParamsMutex.Lock()
	fake.setAdditionalParamsArgsForCall = append(fake.setAdditionalParamsArgsForCall, struct {
		params map[string]string
	}{params})
	fake.setAdditionalParamsMutex.Unlock()
	if fake.SetAdditionalParamsStub != nil {
		fake.SetAdditionalParamsStub(params)
	}
}

func (fake *FakeOauthConsumer) SetAdditionalParamsCallCount() int {
	fake.setAdditionalParamsMutex.RLock()
	defer fake.setAdditionalParamsMutex.RUnlock()
	return len(fake.setAdditionalParamsArgsForCall)
}

func (fake *FakeOauthConsumer) SetAdditionalParamsArgsForCall(i int) map[string]string {
	fake.setAdditionalParamsMutex.RLock()
	defer fake.setAdditionalParamsMutex.RUnlock()
	return fake.setAdditionalParamsArgsForCall[i].params
}

func (fake *FakeOauthConsumer) MakeHttpClient(token *oauth.AccessToken) (*http.Client, error) {
	fake.makeHttpClientMutex.Lock()
	fake.makeHttpClientArgsForCall = append(fake.makeHttpClientArgsForCall, struct {
		token *oauth.AccessToken
	}{token})
	fake.makeHttpClientMutex.Unlock()
	if fake.MakeHttpClientStub != nil {
		return fake.MakeHttpClientStub(token)
	} else {
		return fake.makeHttpClientReturns.result1, fake.makeHttpClientReturns.result2
	}
}

func (fake *FakeOauthConsumer) MakeHttpClientCallCount() int {
	fake.makeHttpClientMutex.RLock()
	defer fake.makeHttpClientMutex.RUnlock()
	return len(fake.makeHttpClientArgsForCall)
}

func (fake *FakeOauthConsumer) MakeHttpClientArgsForCall(i int) *oauth.AccessToken {
	fake.makeHttpClientMutex.RLock()
	defer fake.makeHttpClientMutex.RUnlock()
	return fake.makeHttpClientArgsForCall[i].token
}

func (fake *FakeOauthConsumer) MakeHttpClientReturns(result1 *http.Client, result2 error) {
	fake.MakeHttpClientStub = nil
	fake.makeHttpClientReturns = struct {
		result1 *http.Client
		result2 error
	}{result1, result2}
}

var _ auth.OauthConsumer = new(FakeOauthConsumer)
