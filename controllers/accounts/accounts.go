package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"
	"github.com/alphand/skilltree-server/middleware"
	"github.com/alphand/skilltree-server/models"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type ghCode struct {
	Code string
}

// Handler - Handle All Account Related
type Handler struct {
	Context    context.Context
	OAuth2Conf *oauth2.Config
	dbi        db.DBInvoker
}

//Hello - hello method
func (h *Handler) Hello() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "plain/text")
		io.WriteString(rw, "Hello World! path: "+req.URL.Path)
	}
}

// HandleGithubExchange - exchagne github code to GH token
func (h *Handler) HandleGithubExchange() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		var reqBody ghCode
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqBody)
		if err != nil {
			j, _ := json.Marshal(err)
			rw.Write(j)
			return
		}

		token, err := h.OAuth2Conf.Exchange(h.Context, reqBody.Code)
		if err != nil {
			fmt.Printf("github oauth exchange failed with %s", err)
			j, _ := json.Marshal(err)
			rw.Write(j)
			return
		}

		j, _ := json.Marshal(token)
		rw.Write(j)
	}
}

//RegisterUser - create new user registration
func (h *Handler) RegisterUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		var reqBody models.User
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&reqBody)

		if err != nil {
			j, _ := json.Marshal(err)
			rw.Write(j)
			return
		}

		ctx := req.Context()
		mgoSess := middleware.GetMongoConn(ctx)

		dbStore := h.dbi.NewDataStore(mgoSess, "skilltree-db", "users")
		models.InitUserDBStore(dbStore)

		reqBody.ID = bson.NewObjectId()
		inst, err2 := dbStore.Create(reqBody.ID, reqBody)

		if err2 != nil {
			errJSON, _ := json.Marshal(err2)
			rw.Write(errJSON)
		}

		resp, _ := json.Marshal(inst.(models.User))
		rw.WriteHeader(http.StatusOK)
		rw.Write(resp)
	}
}
