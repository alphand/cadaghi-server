package models

import (
	"time"

	db "github.com/alphand/skilltree-server/database"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collName = "Users"

var (
	coll *mgo.Collection
)

//SetDBStore - set db session for accounts model
func SetDBStore(dbStore *db.DataStore) {
	coll = dbStore.Session.DB(dbStore.DBName).C(collName)
	coll.EnsureIndex(mgo.Index{
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
	CreatedDate time.Time
	UpdatedDate time.Time
}

//Create - create new user
func (u *User) Create() (*User, error) {
	timestamp := time.Now()
	u.ID = bson.NewObjectId()
	u.CreatedDate = timestamp
	u.UpdatedDate = timestamp

	err := coll.Insert(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
