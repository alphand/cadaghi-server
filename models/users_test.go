package models_test

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	db "github.com/alphand/skilltree-server/database"
	"github.com/alphand/skilltree-server/models"

	"github.com/icrowley/fake"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {

	Convey("Given User creation should be validated", t, func() {
		sess := db.InitMongoSession(connStr)
		ds, _ := db.NewMongoStore(sess, dbName, "users")
		models.InitUserDBStore(ds)

		Convey("User can be created", func() {
			timestamp := time.Now().Unix()
			user := &models.User{
				ID:          bson.NewObjectId(),
				FirstName:   fake.FirstName(),
				LastName:    fake.LastName(),
				Email:       fake.EmailAddress(),
				CreatedDate: timestamp,
				UpdatedDate: timestamp,
			}

			err := ds.Create(user)

			So(user.ID, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("User cannot have duplicate email", func() {
			user1 := &models.User{
				ID:        bson.NewObjectId(),
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     fake.EmailAddress(),
			}
			ds.Create(user1)

			user2 := &models.User{
				ID:        bson.NewObjectId(),
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     user1.Email,
			}
			err := ds.Create(user2)

			So(user1.FirstName, ShouldNotEqual, user2.FirstName)
			So(user1.Email, ShouldEqual, user2.Email)
			So(err, ShouldNotBeNil)
		})

		Reset(func() {
			defer sess.Close()
			sess.DB(dbName).DropDatabase()
		})
	})
}
