package models

import (
	db "github.com/alphand/skilltree-server/database"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//InitUserDBStore - set db session for accounts model
func InitUserDBStore(ds db.IDataStore) {
	ds.Coll().EnsureIndex(mgo.Index{
		Name:       "email_unqkey",
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	})
}

// User - user model
type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email       string
	FirstName   string
	LastName    string
	CreatedDate int64
	UpdatedDate int64
}
