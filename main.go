package main

import (
	"fmt"
	"os"

	"github.com/workos/workos-go/v3/pkg/usermanagement"
)

func main() {
	apiKey := os.Getenv("WORKOS_API_KEY")
	clientId := os.Getenv("WORK_OS_CLIENT_ID")
	usermanagement.SetAPIKey(apiKey)

	url, err := usermanagement.GetAuthorizationURL(
		usermanagement.GetAuthorizationURLOpts{
			ClientID:    clientId,
			RedirectURI: "https://testjeddatacenter.azurewebsites.net/workos/callback",
			Provider:    "authkit",
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(url)

}
