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
	sess := m.session.Clone()
	defer sess.Close()
	c := sess.DB(m.DBName).C(m.ColName)
	f(c)
}

// NewMongoStore - Create new MongoDB Handler
func NewMongoStore(connStr, dbName, collName string) (IDataStore, error) {
	m, err := initMongo(connStr)
	if err != nil {
		return nil, err
	}
	m.DBName = dbName
	m.ColName = collName

	return m, nil
}

func initMongo(connStr string) (*MongoDS, error) {
	mongsess, err := mgo.Dial(connStr)
	if err != nil {
		return nil, err
	}

	mongsess.SetMode(mgo.Monotonic, true)
	m := &MongoDS{
		session: mongsess,
	}

	return m, nil
}
