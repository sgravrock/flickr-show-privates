package app

import (
	"fmt"
	"io"
	"github.com/sgravrock/flickr-show-privates/auth"
	"github.com/sgravrock/flickr-show-privates/flickrapi"
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
	httpClient, err := app.authenticator.Authenticate()
	if err != nil {
		fmt.Fprintln(app.stderr, err.Error())
		return 1
	}

	flickrClient := flickrapi.NewClient(httpClient, "https://api.flickr.com")

	fmt.Fprintln(app.stdout, "Downloading photo list")
	photolist, err := flickrClient.GetPhotos(500)
	if err != nil {
		fmt.Fprintln(app.stderr, err.Error())
		return 1
	}

	for i := 0; i < len(photolist); i++ {
		isPublic, err := photolist[i].IsPublic()
		if err != nil {
			fmt.Fprintln(app.stderr, err.Error())
			return 1
		}

		isFamily, err := photolist[i].IsFamily()
		if err != nil {
			fmt.Fprintln(app.stderr, err.Error())
			return 1
		}

		isFriend, err := photolist[i].IsFriend()
		if err != nil {
			fmt.Fprintln(app.stderr, err.Error())
			return 1
		}


		if !(isPublic || isFamily || isFriend) {
			id, err := photolist[i].Id()
			if err != nil {
				fmt.Fprintln(app.stderr, err.Error())
				return 1
			}

			fmt.Println(id)
		}
	}

	fmt.Printf("Got %d photos\n", len(photolist))

	return 0
}