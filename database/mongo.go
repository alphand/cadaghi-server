package db

import mgo "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

type mgoAction func(c *mgo.Collection)

// MongoDS - mongo ds implementation
type MongoDS struct {
	session *mgo.Session
	DBName  string
	ColName string
}

// Create - Mongo Create Record function
func (m *MongoDS) Create(i interface{}) (err error) {
	m.execMgoAction(func(c *mgo.Collection) {
		err = c.Insert(i)
	})
	return
}

//GetByID - Mongo get record by ID
func (m *MongoDS) GetByID(id string) (o interface{}, err error) {
	m.execMgoAction(func(c *mgo.Collection) {
		err = c.FindId(bson.ObjectIdHex(id)).One(&o)
	})
	return
}

//GetAll - Mongo get all record
func (m *MongoDS) GetAll(i interface{}) (o []interface{}, err error) {
	m.execMgoAction(func(c *mgo.Collection) {
		if i == nil {
			i = &bson.M{}
		}
		err = c.Find(i).All(&o)
	})
	return
}

//SetIndex - set index for collection
func (m *MongoDS) SetIndex(i interface{}) (err error) {
	m.execMgoAction(func(c *mgo.Collection) {
		err = c.EnsureIndex(i.(mgo.Index))
	})
	return
}

func (m *MongoDS) execMgoAction(f mgoAction) {
	sess := m.session.Copy()
	defer sess.Close()
	c := sess.DB(m.DBName).C(m.ColName)
	f(c)
}

// NewMongoStore - Create new MongoDB Handler
func NewMongoStore(session *mgo.Session, dbName, collName string) (IDataStore, error) {
	m := initMongo(session)

	m.DBName = dbName
	m.ColName = collName

	return m, nil
}

func initMongo(session *mgo.Session) *MongoDS {
	return &MongoDS{
		session: session,
	}
}

//InitMongoSession - initialize mongo connection
func InitMongoSession(connStr string) *mgo.Session {
	mongsess, err := mgo.Dial(connStr)
	if err != nil {
		panic(err)
	}

	mongsess.SetMode(mgo.Monotonic, true)
	return mongsess
}
