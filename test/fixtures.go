package fixtures

import (
	"time"

	"github.com/alphand/skilltree-server/models"
	"github.com/icrowley/fake"
	"gopkg.in/mgo.v2/bson"
)

//NewUser - Create fixture for new user
func NewUser() *models.User {
	timestamp := time.Now().Unix()
	return &models.User{
		ID:          bson.NewObjectId(),
		FirstName:   fake.FirstName(),
		LastName:    fake.LastName(),
		Email:       fake.EmailAddress(),
		CreatedDate: timestamp,
		UpdatedDate: timestamp,
	}
}

//NewIntegration - Create fixture for new integration
func NewIntegration(user *models.User, provider string) *models.UserIntegration {
	timestamp := time.Now().Unix()

	return &models.UserIntegration{
		ID:          bson.NewObjectId(),
		UserID:      user.ID,
		Provider:    provider,
		AccessToken: bson.NewObjectId().String(),
		ExpireIn:    time.Now().AddDate(0, 1, 0).Unix(),
		CreatedDate: timestamp,
		UpdatedDate: timestamp,
	}
}
