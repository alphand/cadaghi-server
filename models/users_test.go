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

		sess, _ := db.NewSession("192.168.18.129")
		dbStore := db.NewDataStore(sess, "testusers", "users")
		models.InitUserDBStore(dbStore)

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

			_, err := dbStore.Create(user.ID, user)

			So(user.ID, ShouldNotBeEmpty)
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
			dbStore.Create(user1.ID, user1)

			user2 := &models.User{
				ID:        bson.NewObjectId(),
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     user1.Email,
			}
			_, err := dbStore.Create(user2.ID, user2)

			So(user1.FirstName, ShouldNotEqual, user2.FirstName)
			So(user1.Email, ShouldEqual, user2.Email)
			So(err, ShouldNotBeNil)
		})

		Reset(func() {
			sess.DB(dbStore.DBName).DropDatabase()
		})
	})
}
