package datastore

import (
	"encoding/json"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeLog(action string, data interface{}, usernName string) ([]Log, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	logRaw := []Log{
		Log{
			Datetime:   primitive.DateTime(millis.NowInMillis()),
			Action:     action,
			Value:      string(b),
			ModifiedBy: usernName,
		},
	}
	return logRaw, nil
}

func makeLogRaw(action string, data interface{}, userName string) (*Log, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	logRaw := &Log{
		Datetime:   primitive.DateTime(millis.NowInMillis()),
		Action:     action,
		Value:      string(b),
		ModifiedBy: userName,
	}

	return logRaw, nil
}

func makeHistoryLogRaw(filename, url, userName string) *IformHistoryLog {

	logRaw := &IformHistoryLog{
		CreatedDate: primitive.DateTime(millis.NowInMillis()),
		UpdatedDate: primitive.DateTime(millis.NowInMillis()),
		Url:         url,
		fileName:    filename,
		ModifiedBy:  userName,
	}

	return logRaw
}
