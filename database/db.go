package db

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IMongoSess - monggo session interface
type IMongoSess interface {
	DB(name string) IMongoDatabase
	Copy() *mgo.Session
	Close()
}

// IMongoDatabase - mongo db interface
type IMongoDatabase interface {
	C(name string) *mgo.Collection
	DropDatabase()
}

//IDataStore - DataStore interface for DB Ops
type IDataStore interface {
	Coll() *mgo.Collection
	Create(id bson.ObjectId, o interface{}) (interface{}, error)
	// GetById(id string) (interface{}, error)
	// Put(id string) (interface{}, error)
	// Delete(id string, soft bool)
}

// IDBStore - DBStore interface
type IDBStore interface {
	NewDataStore(...interface{}) IDataStore
	NewSession(...interface{}) (IMongoSess, error)
}

// DBInvoker - Real DB Creation
type DBInvoker struct {
}

// NewDataStore - Create new DB Handler
func (d *DBInvoker) NewDataStore(session IMongoSess, dbName, collName string) *DataStore {
	return &DataStore{
		session:  session,
		DBName:   dbName,
		CollName: collName,
	}
}

// NewSession - crate new mongo session
func (d *DBInvoker) NewSession(connStr string) (*mgo.Session, error) {
	mongsess, err := mgo.Dial(connStr)
	if err != nil {
		return nil, err
	}

	return mongsess, nil
}

//DataStore - handle DB Connection
type DataStore struct {
	session  IMongoSess
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
