package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	acc "github.com/alphand/skilltree-server/accounts"
	"github.com/alphand/skilltree-server/middleware"
	ghoauth "golang.org/x/oauth2/github"
)

var (
	port string

	oauthGithubConf = &oauth2.Config{
		ClientID:     "6a418b2d916f57eb921c",
		ClientSecret: "35d7548d80049e6b4ac962e77f81e095330702b8",
		Scopes:       []string{"user:email"},
		Endpoint:     ghoauth.Endpoint,
	}
)

func init() {
	flag.StringVar(&port, "p", ":3000", "Webserver port address")
}

func main() {
	flag.Parse()

	neg := SetupNegroni()

	server := &http.Server{
		Addr:    port,
		Handler: neg,
	}

	log.Println("Webserver is ready at: " + port)
	log.Fatal(server.ListenAndServe())
}

// SetupNegroni - Setup server initialization for testing
func SetupNegroni() *negroni.Negroni {
	router := mux.NewRouter()

	neg := negroni.New()
	neg.Use(negroni.NewRecovery())
	neg.Use(middleware.NewContext(&middleware.UUIDGen{}))

	accHdl := &acc.Handler{
		Context:    oauth2.NoContext,
		OAuth2Conf: oauthGithubConf,
	}

	router.HandleFunc("/", accHdl.Hello())
	router.HandleFunc("/accounts/github", accHdl.HandleGithubExchange()).Methods("POST")

	neg.UseHandler(router)

	return neg
}

func helloFunc(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World! path: "+req.URL.Path)
}
