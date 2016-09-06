package db_test

import (
	"testing"

	"github.com/alphand/skilltree-server/database"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabase(t *testing.T) {
	Convey("Given database is ready for testing", t, func() {
		sess, _ := db.NewSession("192.168.18.129")
		defer sess.Close()

		ds := db.NewDB(sess, "test")

		Convey("DB is ready to be used", func() {
			So(ds.Session, ShouldNotBeEmpty)
			So(ds.DBName, ShouldEqual, "test")
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

	})
}
