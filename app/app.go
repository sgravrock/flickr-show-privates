package app

import (
	"fmt"
	"github.com/sgravrock/flickr-show-privates/auth"
	"github.com/sgravrock/flickr-show-privates/flickrapi"
	"io"
	"io/ioutil"
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

	outfile, err := ioutil.TempFile("", "photos.*.html")
	if err != nil {
		return err
	}

	defer func() {
		cerr := outfile.Close()
		if err == nil {
			err = cerr
		}
	}()

	fmt.Fprintf(app.stdout, "Writing to %s\n", outfile.Name())
	writer := newErrorTrackingWriter(outfile)
	writer.println("<html><body>")

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

			writer.printf("%s<br>", id)
		}
	}

	writer.println("</body></html>")
	// Note: assignment to err before return is load-bearing.
	// See the deferred closure above.
	err = writer.err
	return err
}

type errorTrackingWriter struct {
	file io.Writer
	err  error
}

func newErrorTrackingWriter(f io.Writer) errorTrackingWriter {
	return errorTrackingWriter{
		file: f,
		err: nil,
	}
}

func (w *errorTrackingWriter) println(s string) {
	if w.err == nil {
		_, w.err = fmt.Fprintln(w.file, s)
	}
}

func (w *errorTrackingWriter) printf(s string, a ...interface{}) {
	if w.err == nil {
		fmt.Fprintf(w.file, s, a...)
	}
}