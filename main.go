package main

import (
	"fmt"
	"os"

	"github.com/sgravrock/flickr-show-privates/app"
	"github.com/sgravrock/flickr-show-privates/auth"
)

func main() {
	key := requireEnv("FLICKR_API_KEY")
	secret := requireEnv("FLICKR_API_SECRET")
	authenticator := auth.NewAuthenticator(key, secret, nil, nil)
	exitcode := app.Run("https://api.flickr.com", authenticator,
		os.Stdout, os.Stderr)
	os.Exit(exitcode)
}

func requireEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Fprintf(os.Stderr, "Please set the %s environment variable\n", name)
		os.Exit(1)
	}
	return value
}