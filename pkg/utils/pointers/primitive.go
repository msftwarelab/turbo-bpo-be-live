package pointers

import (
	"errors"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrimitiveDateTime(t *time.Time) *primitive.DateTime {
	ct := time.Now()

	priTime := new(primitive.DateTime)
	if t != nil {
		*priTime = primitive.DateTime(millis.InMillis(*t))
		return priTime
	}
	*priTime = primitive.DateTime(millis.InMillis(ct))
	return priTime

}

func PrimitiveDateTimeAddHr(hr float64) *primitive.DateTime {
	ct := time.Now().Add(time.Hour * time.Duration(hr))
	priTime := new(primitive.DateTime)
	*priTime = primitive.DateTime(millis.InMillis(ct))
	return priTime

}
func PrimitiveDateTimeAddHrInt(hr int) *primitive.DateTime {
	ct := time.Now().Add(time.Hour * time.Duration(hr))
	priTime := new(primitive.DateTime)
	*priTime = primitive.DateTime(millis.InMillis(ct))
	return priTime

}

func PrimitiveDateTimeToInt64(t *primitive.DateTime) *int64 {
	if t != nil {
		return Int64(int64(*t))
	}
	return nil

}

// Time returns the date as a time type.
func PrimativeToDateTime(d primitive.DateTime) time.Time {
	return time.Unix(int64(d)/1000, int64(d)%1000*1000000)
}

//Convert string time to primative datetime format
func StringTimeToPrimitive(s string) (*primitive.DateTime, error) {
	layout := "2006-01-02"
	datetime, err := time.Parse(layout, s)
	if err != nil {
		return nil, errors.New("invalid time format")
	}
	return PrimitiveDateTime(&datetime), nil

}

//Convert string datetime to primative datetime format
func StringTDateimeToPrimitive(s string) (*primitive.DateTime, error) {
	layout := time.RFC3339
	datetime, err := time.Parse(layout, s)
	if err != nil {
		return nil, errors.New("invalid time format")
	}
	return PrimitiveDateTime(&datetime), nil

}

//Convert string time to primative datetime format and add Hr
func StringTimeToPrimitiveAddHr(s string, addHr float32) (*primitive.DateTime, error) {
	layout := "2006-01-02"
	datetime, err := time.Parse(layout, s)
	if err != nil {
		return nil, errors.New("invalid time format")
	}

	//duedatetime := pointers.PrimativeToDateTime(*v.DueDateTime).Add(time.Hour * time.Duration(48))
	datetime = datetime.Add(time.Hour * time.Duration(addHr))
	return PrimitiveDateTime(&datetime), nil

}
