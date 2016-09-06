package models_test

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"
	"github.com/alphand/skilltree-server/models"
	"github.com/alphand/skilltree-server/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserIntegration(t *testing.T) {
	Convey("Given DB is setup properly", t, func() {
		sess, _ := db.NewSession("192.168.18.129")
		userStore := db.NewDataStore(sess, "testusers", "users")
		userIntegrationStore := db.NewDataStore(sess, "testusers", "userintegrations")
		models.InitUserDBStore(userStore)
		models.InitUserIntegrationDBStore(userIntegrationStore)

		user := fixtures.NewUser()

		userStore.Create(user)

		Convey("User can be integrated", func() {
			timestamp := time.Now().Unix()
			ghintegration := &models.UserIntegration{
				ID:          bson.NewObjectId(),
				UserID:      user.ID,
				Provider:    "github",
				AccessToken: bson.NewObjectId().String(),
				ExpireIn:    time.Now().AddDate(0, 1, 0).Unix(),
				CreatedDate: timestamp,
				UpdatedDate: timestamp,
			}

			_, err := userIntegrationStore.Create(ghintegration)

			So(ghintegration.ID, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("User cannot have duplicate integration", func() {
			intg1 := fixtures.NewIntegration(user, "github")
			userIntegrationStore.Create(intg1)

			intg2 := fixtures.NewIntegration(user, "github")
			_, err := userIntegrationStore.Create(intg2)

			So(err, ShouldNotBeNil)
		})

		Reset(func() {
			userStore.DropDB()
			userIntegrationStore.DropDB()
			sess.Close()
		})
	})
}
