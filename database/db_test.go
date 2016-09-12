package db_test

import (
	"log"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabase(t *testing.T) {
	Convey("Given database is ready for testing", t, func() {
		sess, _ := db.NewSession("192.168.18.129")

		ds := db.NewDataStore(sess, "test", "users")

		Convey("DB is ready to be used", func() {
			So(ds.Session, ShouldNotBeEmpty)
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

			ds.Create(rec)

			var res []recd
			_ = ds.C().Find(bson.M{}).All(&res)

			So(len(res), ShouldBeGreaterThanOrEqualTo, 1)
		})

		Convey("DB can dropdb", func() {
			type recd struct {
				ID   bson.ObjectId `bson:"_id,omitempty"`
				Name string
			}

			rec := &recd{
				ID:   bson.NewObjectId(),
				Name: "James",
			}

			ds.Create(rec)

			ds.DropDB()

			names, _ := ds.Session.DatabaseNames()
			log.Println("dbnames: ", names)

			So(len(names), ShouldEqual, 1)
		})

		// Convey("DB is dropped", func() {
		// 	ds.Session.DB(ds.DBName).C("users")

		// 	var dbFound = make([]string, 0, 2)
		// 	names, _ := ds.Session.DatabaseNames()
		// 	log.Println("dbfound", dbFound, len(dbFound), names)

		// 	So(len(names), ShouldEqual, 2)

		// 	ds.DropDB()

		// 	names, _ = ds.Session.DatabaseNames()
		// 	for _, v := range names {
		// 		if v == ds.DBName {
		// 			dbFound = append(dbFound, v)
		// 			break
		// 		}
		// 	}

		// 	So(len(dbFound), ShouldEqual, 1)
		// })

		Reset(func() {
			sess.Close()
		})

	})
}
