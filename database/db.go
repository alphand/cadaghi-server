package db

import mgo "gopkg.in/mgo.v2"

//DataStore - handle DB Connection
type DataStore struct {
	Session *mgo.Session
	DBName  string
}

//DropDB - to drop db for testing purposes
func (d *DataStore) DropDB() {
	newSess := d.Session.Clone()
	newSess.DB(d.DBName).DropDatabase()
}

// NewDB - Create new DB Connection
func NewDB(session *mgo.Session, dbName string) *DataStore {
	return &DataStore{
		Session: session,
		DBName:  dbName,
	}
}

// NewSession - crate new mongo session
func NewSession(connStr string) (*mgo.Session, error) {
	return mgo.Dial(connStr)
}
