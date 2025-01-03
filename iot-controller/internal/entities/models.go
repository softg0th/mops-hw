package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	bsonPrimitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjectID = bsonPrimitive.ObjectID

type Document struct {
	ID              ObjectID
	DeviceID        int
	Timestamp       time.Time
	SomeUsefulField int
}

func NewDocument(deviceID int32, timestamp time.Time, someUsefulField int32) *Document {
	return &Document{
		ID:              primitive.NewObjectID(),
		DeviceID:        int(deviceID),
		Timestamp:       timestamp,
		SomeUsefulField: int(someUsefulField),
	}
}
