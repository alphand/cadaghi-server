package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthGithubConf = &oauth2.Config{
		ClientID:     "6a418b2d916f57eb921c",
		ClientSecret: "35d7548d80049e6b4ac962e77f81e095330702b8",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
	oauthStateString = "SuperWierdandLongTextTOdeterminethis"
)

type githubToken struct {
	accessToken string
}

type ghCode struct {
	Code string
}

// HandleGithubExchange - exchagne github code to GH token
func HandleGithubExchange(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var reqBody ghCode
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		j, _ := json.Marshal(err)
		rw.Write(j)
		return
	}

	token, err := oauthGithubConf.Exchange(oauth2.NoContext, reqBody.Code)
	if err != nil {
		fmt.Printf("github oauth exchange failed with %s", err)
		j, _ := json.Marshal(err)
		rw.Write(j)
		return
	}

	j, _ := json.Marshal(token)
	rw.Write(j)
}
