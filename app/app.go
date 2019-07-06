package app

import (
	"fmt"
	"io"
	"github.com/sgravrock/flickr-show-privates/auth"
	"github.com/sgravrock/flickr-show-privates/flickrapi"
)

func Run(baseUrl string, authenticator auth.Authenticator,
	stdout io.Writer, stderr io.Writer) error {

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

func (app *application) Run() error {
	httpClient, err := app.authenticator.Authenticate()
	if err != nil {
		return err
	}

	flickrClient := flickrapi.NewClient(httpClient, "https://api.flickr.com")

	fmt.Fprintln(app.stdout, "Downloading photo list")
	photolist, err := flickrClient.GetPhotos(500)
	if err != nil {
		return err
	}

	for i := 0; i < len(photolist); i++ {
		isPublic, err := photolist[i].IsPublic()
		if err != nil {
			return err
		}

		isFamily, err := photolist[i].IsFamily()
		if err != nil {
			return err
		}

		isFriend, err := photolist[i].IsFriend()
		if err != nil {
			return err
		}

		if !(isPublic || isFamily || isFriend) {
			id, err := photolist[i].Id()
			if err != nil {
				return err
			}

			fmt.Println(id)
		}
	}

	return nil
}