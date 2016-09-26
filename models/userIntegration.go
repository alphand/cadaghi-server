package models

import (
	db "github.com/alphand/skilltree-server/database"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//InitUserIntegrationDBStore - set db session for accounts model
func InitUserIntegrationDBStore(ds db.IDataStore) {
	ds.SetIndex(mgo.Index{
		Name:       "idx_user_provider",
		Key:        []string{"userid", "provider"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	})
}

//UserIntegration - collection of user integration
type UserIntegration struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      bson.ObjectId
	Provider    string
	AccessToken string
	ExpireIn    int64
	CreatedDate int64
	UpdatedDate int64
}
