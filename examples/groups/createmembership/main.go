package main

import (
	"flag"
	"fmt"

	"github.com/kylelemons/godebug/pretty"
	"github.com/ttacon/box"
	"golang.org/x/oauth2"
)

var (
	clientId     = flag.String("cid", "", "OAuth Client ID")
	clientSecret = flag.String("csec", "", "OAuth Client Secret")

	accessToken  = flag.String("atok", "", "Access Token")
	refreshToken = flag.String("rtok", "", "Refresh Token")

	userID  = flag.String("uid", "", "user to create membership for")
	groupID = flag.String("gid", "", "group to add user to")
	role    = flag.String("role", "", "role for user to have")
)

func main() {
	flag.Parse()

	if len(*clientId) == 0 || len(*clientSecret) == 0 ||
		len(*accessToken) == 0 || len(*refreshToken) == 0 ||
		len(*userID) == 0 || len(*groupID) == 0 {
		fmt.Println("unfortunately all flags must be provided")
		return
	}

	// Set our OAuth2 configuration up
	var (
		configSource = box.NewConfigSource(
			&oauth2.Config{
				ClientID:     *clientId,
				ClientSecret: *clientSecret,
				Scopes:       nil,
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://app.box.com/api/oauth2/authorize",
					TokenURL: "https://app.box.com/api/oauth2/token",
				},
				RedirectURL: "http://localhost:8080/handle",
			},
		)
		tok = &oauth2.Token{
			TokenType:    "Bearer",
			AccessToken:  *accessToken,
			RefreshToken: *refreshToken,
		}
		c = configSource.NewClient(tok)
	)

	resp, membershipEntry, err := c.GroupService().AddUserToGroup(*userID, *groupID, *role)
	fmt.Println("resp: ", resp)
	fmt.Println("err: ", err)
	pretty.Println(membershipEntry)
	// Print out the new tokens for next time
	fmt.Printf("\n%#v\n", tok)
}
