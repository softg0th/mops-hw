package entities

import (
	bsonPrimitive "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ObjectID = bsonPrimitive.ObjectID

type Document struct {
	ID              ObjectID
	deviceID        int
	timestamp       time.Time
	someUsefulField int
}
