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

	streamPosition = flag.String("pos", "0", "position in stream to query")
	streamType     = flag.String("type", "all", "types of events to query")
	limit          = flag.Int("lim", 100, "limit of events to query")
)

func main() {
	flag.Parse()

	if len(*clientId) == 0 || len(*clientSecret) == 0 ||
		len(*accessToken) == 0 || len(*refreshToken) == 0 {
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

	resp, eventCollection, err := c.EventService().Events(box.EventQueryOptions{
		StreamPosition: *streamPosition,
		StreamType:     *streamType,
		Limit:          *limit,
	})
	fmt.Println("resp: ", resp)
	fmt.Println("err: ", err)
	pretty.Print(eventCollection)
	fmt.Println("\n# events: ", len(eventCollection.Entries))

	// Print out the new tokens for next time
	fmt.Printf("\n%#v\n", tok)
}
