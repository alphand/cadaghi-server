package db

import mgo "gopkg.in/mgo.v2"

//DataStore - handle DB Connection
type DataStore struct {
	Session  *mgo.Session
	DBName   string
	CollName string
}

//DropDB - to drop db for testing purposes
func (d *DataStore) DropDB() {
	d.Session.DB(d.DBName).DropDatabase()
}

//C - DataStore Collection
func (d *DataStore) C() *mgo.Collection {
	return d.Session.DB(d.DBName).C(d.CollName)
}

//Create - create record
func (d *DataStore) Create(v interface{}) (interface{}, error) {
	err := d.C().Insert(v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// NewDataStore - Create new DB Connection
func NewDataStore(session *mgo.Session, dbName, collName string) *DataStore {
	return &DataStore{
		Session:  session,
		DBName:   dbName,
		CollName: collName,
	}
}

// NewSession - crate new mongo session
func NewSession(connStr string) (*mgo.Session, error) {
	return mgo.Dial(connStr)
}
