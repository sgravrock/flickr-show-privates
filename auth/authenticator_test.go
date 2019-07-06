package auth_test

import (
	"errors"
	"net/http"
	. "github.com/sgravrock/flickr-show-privates/auth"

	"github.com/mrjones/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sgravrock/flickr-show-privates/auth/authfakes"
)

var _ = Describe("Authenticate", func() {
	var subject Authenticator
	var oauthClient authfakes.FakeOauthClient
	var oauthConsumer *authfakes.FakeOauthConsumer
	var ui *authfakes.FakeUiAdapter

	BeforeEach(func() {
		oauthClient = *new(authfakes.FakeOauthClient)
		oauthConsumer = new(authfakes.FakeOauthConsumer)
		ui = new(authfakes.FakeUiAdapter)
		oauthClient.NewConsumerReturns(oauthConsumer)
		subject = NewAuthenticator("token", "secret", &oauthClient, ui)
	})

	Context("When obtaining a request token fails", func() {
		var tokenError error

		BeforeEach(func() {
			tokenError = errors.New("nope")
			oauthConsumer.GetRequestTokenAndUrlReturns(nil, "", tokenError)
		})

		It("returns the error", func() {
			t, err := subject.Authenticate()
			Expect(t).To(BeNil())
			Expect(err).To(BeIdenticalTo(tokenError))
		})
	})

	Context("When a request token is obtained", func() {
		var requestToken oauth.RequestToken

		BeforeEach(func() {
			requestToken = oauth.RequestToken{"the request token", ""}
			oauthConsumer.GetRequestTokenAndUrlReturns(&requestToken,
				"the://url", nil)
		})

		It("prompts the user for authentication", func() {
			subject.Authenticate()
			Expect(ui.PromptForAccessCodeCallCount()).To(Equal(1))
			Expect(ui.PromptForAccessCodeArgsForCall(0)).To(Equal("the://url"))
		})

		Context("When obtaining an access code fails", func() {
			var codeError error

			BeforeEach(func() {
				codeError = errors.New("nope")
				ui.PromptForAccessCodeReturns("", codeError)
			})

			It("returns the error", func() {
				t, err := subject.Authenticate()
				Expect(t).To(BeNil())
				Expect(err).To(BeIdenticalTo(codeError))
			})
		})

		Context("When the user enters an access code", func() {
			BeforeEach(func() {
				ui.PromptForAccessCodeReturns("the code", nil)
			})

			It("authorizes the request token and code", func() {
				subject.Authenticate()
				Expect(oauthConsumer.AuthorizeTokenCallCount()).To(Equal(1))
				token, code := oauthConsumer.AuthorizeTokenArgsForCall(0)
				Expect(token).To(Equal(&requestToken))
				Expect(code).To(Equal("the code"))
			})

			Context("When authorization fails", func() {
				var authError error

				BeforeEach(func() {
					authError = errors.New("nope")
					oauthConsumer.AuthorizeTokenReturns(nil, authError)
				})

				It("returns the error", func() {
					t, err := subject.Authenticate()
					Expect(t).To(BeNil())
					Expect(err).To(BeIdenticalTo(authError))
				})

			})

			Context("When authorization succeeds", func() {
				var accessToken oauth.AccessToken
				var httpClient http.Client

				BeforeEach(func() {
					accessToken = oauth.AccessToken{
						"access token", "", nil,
					}
					oauthConsumer.AuthorizeTokenReturns(&accessToken, nil)
					httpClient = http.Client{}
					oauthConsumer.MakeHttpClientReturns(&httpClient, nil)
				})

				It("returns the created HTTP client", func() {
					result, err := subject.Authenticate()
					Expect(result).To(BeIdenticalTo(&httpClient))
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
