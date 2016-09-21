package db_test

import (
	"log"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabase(t *testing.T) {
	Convey("Given mongo database is ready for testing", t, func() {
		connStr := "192.168.18.129"
		dbName := "test"
		mgodb, err := db.NewMongoStore(connStr, dbName, "users")

		Convey("DB is ready to be used", func() {
			So(err, ShouldBeNil)
		})

		Convey("DB can insert record", func() {
			type recd struct {
				ID   bson.ObjectId `bson:"_id,omitempty"`
				Name string
			}

			rec := &recd{
				ID:   bson.NewObjectId(),
				Name: "James",
			}

			err2 := mgodb.Create(rec)

			// var res []recd
			res := make([]recd, 0)
			mgodb.GetAll(bson.M{}, res)

			log.Println("res", res)

			So(err2, ShouldBeNil)
			// So(len(res), ShouldBeGreaterThanOrEqualTo, 1)
		})

		Reset(func() {
			sess, _ := mgo.Dial(connStr)
			defer sess.Close()

			sess.DB(dbName).DropDatabase()
		})

	})
}
