package models_test

import (
	"testing"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/alphand/skilltree-server/database"
	"github.com/alphand/skilltree-server/models"
	"github.com/alphand/skilltree-server/test"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	connStr = "192.168.18.129"
	dbName  = "testapp"
)

func TestUserIntegration(t *testing.T) {
	Convey("Given DB is setup properly", t, func() {
		dsUser, _ := db.NewMongoStore(connStr, dbName, "users")
		dsIntUser, _ := db.NewMongoStore(connStr, dbName, "userintegrations")

		models.InitUserDBStore(dsUser)
		models.InitUserIntegrationDBStore(dsIntUser)

		user := fixtures.NewUser()

		dsUser.Create(user)

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

			err := dsIntUser.Create(ghintegration)

			So(ghintegration.ID, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("User cannot have duplicate integration", func() {
			intg1 := fixtures.NewIntegration(user, "github")
			dsIntUser.Create(intg1)

			intg2 := fixtures.NewIntegration(user, "github")
			err := dsIntUser.Create(intg2)

			So(err, ShouldNotBeNil)
		})

		Reset(func() {
			sess, _ := mgo.Dial(connStr)
			defer sess.Close()
			sess.DB(dbName).DropDatabase()
		})
	})
}
