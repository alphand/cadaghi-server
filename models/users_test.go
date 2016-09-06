package models_test

import (
	"testing"

	"github.com/alphand/skilltree-server/database"
	"github.com/alphand/skilltree-server/models"

	"github.com/icrowley/fake"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("Given DB is setup properly", t, func() {
		sess, _ := db.NewSession("192.168.18.129")
		dbStore := db.NewDB(sess, "testusers")
		models.SetDBStore(dbStore)

		Convey("User can be created", func() {
			user := &models.User{
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     fake.EmailAddress(),
			}

			_, err := user.Create()

			So(user.ID, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("User cannot have duplicate email", func() {
			user1 := &models.User{
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     fake.EmailAddress(),
			}
			user1.Create()

			user2 := &models.User{
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
				Email:     user1.Email,
			}
			_, err := user2.Create()

			So(user1.FirstName, ShouldNotEqual, user2.FirstName)
			So(user1.Email, ShouldEqual, user2.Email)
			So(err, ShouldNotBeNil)
		})

		Reset(func() {
			dbStore.DropDB()
			sess.Close()
		})
	})
}
