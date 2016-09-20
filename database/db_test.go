package db_test

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabase(t *testing.T) {
	Convey("Given database is ready for testing", t, func() {
		dbs := db.DBInvoker{}

		sess, _ := dbs.NewSession("192.168.18.129")

		ds := dbs.NewDataStore(sess, "test", "users")

		Convey("DB is ready to be used", func() {
			So(ds.Coll(), ShouldNotBeNil)
			So(ds.DBName, ShouldEqual, "test")
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

			ds.Create(rec.ID, rec)

			var res []recd
			_ = ds.Coll().Find(bson.M{}).All(&res)

			So(len(res), ShouldBeGreaterThanOrEqualTo, 1)
		})

		Reset(func() {
			sess.DB(ds.DBName).DropDatabase()
			sess.Close()
		})

	})
}
