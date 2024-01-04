package datastore

import (
	sysStrings "strings"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string              `bson:"userID,omitempty"`
	RecordType      string              `bson:"recordType,omitempty"`
	Company         string              `bson:"company,omitempty"`
	WebSite         string              `bson:"webSite,omitempty"`
	Username        string              `bson:"username,omitempty"`
	Password        string              `bson:"password,omitempty"`
	Question1       *string             `json:",omitempty" bson:"question1,omitempty"`
	Answer1         *string             `json:",omitempty" bson:"answer1,omitempty"`
	Question2       *string             `json:",omitempty" bson:"question2,omitempty"`
	Answer2         *string             `bson:"answer2,omitempty"`
	Question3       *string             `bson:"question3,omitempty"`
	Answer3         *string             `bson:"answer3,omitempty"`
	Others          *string             `bson:"others,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
	Logs            []Log               `bson:"logs,omitempty"`
}

func (u *Account) ToModels() *models.Account {

	logList := make([]*models.Log, 0)
	for _, v := range u.Logs {
		logList = append(logList, v.ToModels())
	}

	return &models.Account{
		ID:              u.ID.Hex(),
		RecordType:      sysStrings.ToUpper(u.RecordType),
		Company:         u.Company,
		WebSite:         u.WebSite,
		Username:        u.Username,
		Password:        u.Password,
		Question1:       u.Question1,
		Answer1:         u.Answer1,
		Question2:       u.Question2,
		Answer2:         u.Answer2,
		Question3:       u.Question3,
		Answer3:         u.Answer3,
		Others:          u.Others,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		LastUpdateTime:  strings.ToObject(TimeConversion(u.LastUpdateTime)),
		Logs:            logList,
	}
}
