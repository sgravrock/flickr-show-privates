package app

import (
	"fmt"
	"io"
	"github.com/sgravrock/flickr-show-privates/auth"
)

func Run(baseUrl string, authenticator auth.Authenticator,
	stdout io.Writer, stderr io.Writer) int {

	ftg := application{
		baseUrl:       baseUrl,
		authenticator: authenticator,
		stdout:        stdout,
		stderr:        stderr,
	}
	return ftg.Run()
}

type application struct {
	baseUrl       string
	authenticator auth.Authenticator
	stdout        io.Writer
	stderr        io.Writer
}

func (app *application) Run() int {
	_, err := app.authenticator.Authenticate()
	if err != nil {
		fmt.Fprintln(app.stderr, err.Error())
		return 1
	}

	return 0
}