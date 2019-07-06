package app_test

import (
	"bytes"
	"net/http"

	. "github.com/sgravrock/flickr-show-privates/app"
	"github.com/sgravrock/flickr-show-privates/auth/authfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"errors"
)

var _ = Describe("App", func() {
	var (
		authenticator *authfakes.FakeAuthenticator
		ts            *ghttp.Server
		stdout        *bytes.Buffer
		stderr        *bytes.Buffer
		err           error
	)

	BeforeEach(func() {
		authenticator = new(authfakes.FakeAuthenticator)
		authenticator.AuthenticateReturns(new(http.Client), nil)
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		ts = ghttp.NewServer()
		ts.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/services/rest/",
					"format=json&method=flickr.test.login&nojsoncallback=1"),
				ghttp.RespondWith(http.StatusNotFound, "nope"),
			),
		)
	})

	JustBeforeEach(func() {
		err = Run(ts.URL(), authenticator, stdout, stderr)
	})

	It("authenticates the user", func() {
		Expect(authenticator.AuthenticateCallCount()).To(Equal(1))
	})

	Describe("When authentication fails", func() {
		BeforeEach(func() {
			authenticator.AuthenticateReturns(nil, errors.New("nope"))
		})

		It("returns an error", func() {
			Expect(err).NotTo(BeNil())
		})
	})
})
