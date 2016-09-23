package db_test

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"
	"github.com/icrowley/fake"

	. "github.com/smartystreets/goconvey/convey"
)

type recd struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
}

func TestDatabase(t *testing.T) {
	Convey("Given mongo database is ready for testing", t, func() {
		connStr := "192.168.18.129"
		dbName := "test"
		mgodb, err := db.NewMongoStore(connStr, dbName, "users")

		Convey("DB is ready to be used", func() {
			So(err, ShouldBeNil)
		})

		Convey("DB can insert & retrieve a record", func() {

			rec := &recd{
				ID:   bson.NewObjectId(),
				Name: "James",
			}

			err2 := mgodb.Create(rec)

			var r recd
			o, _ := mgodb.GetByID(rec.ID.Hex())

			v, _ := bson.Marshal(o)
			bson.Unmarshal(v, &r)

			So(err2, ShouldBeNil)
			So(r.Name, ShouldEqual, "James")
			So(r.ID.Hex(), ShouldEqual, rec.ID.Hex())
		})

		Convey("DB can return multiple rows", func() {
			rec1 := &recd{
				ID:   bson.NewObjectId(),
				Name: fake.FirstName(),
			}

			rec2 := &recd{
				ID:   bson.NewObjectId(),
				Name: fake.FirstName(),
			}

			errIns := mgodb.Create(rec1)
			errIns = mgodb.Create(rec2)

			o, _ := mgodb.GetAll(nil)

			var res []recd
			for _, dt := range o {
				v, _ := bson.Marshal(dt)
				var sgl recd
				bson.Unmarshal(v, &sgl)
				res = append(res, sgl)
			}

			So(errIns, ShouldBeNil)
			So(len(res), ShouldEqual, 2)
		})

		Reset(func() {
			sess, _ := mgo.Dial(connStr)
			defer sess.Close()
			sess.DB(dbName).DropDatabase()
		})

	})
}
