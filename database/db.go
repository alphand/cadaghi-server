package db

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//IDataStore - DataStore interface for DB Ops
type IDataStore interface {
	Coll() *mgo.Collection
	Create(id bson.ObjectId, o interface{}) (interface{}, error)
	// GetById(id string) (interface{}, error)
	// Put(id string) (interface{}, error)
	// Delete(id string, soft bool)
}

//DataStore - handle DB Connection
type DataStore struct {
	session  *mgo.Session
	DBName   string
	CollName string
}

//Coll - DataStore Collection
func (d *DataStore) Coll() *mgo.Collection {
	return d.session.DB(d.DBName).C(d.CollName)
}

//Create - create record
func (d *DataStore) Create(id bson.ObjectId, v interface{}) (interface{}, error) {
	_, err := d.Coll().UpsertId(id, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// NewDataStore - Create new DB Connection
func NewDataStore(session *mgo.Session, dbName, collName string) *DataStore {
	return &DataStore{
		session:  session,
		DBName:   dbName,
		CollName: collName,
	}
}

// NewSession - crate new mongo session
func NewSession(connStr string) (*mgo.Session, error) {
	return mgo.Dial(connStr)
}
