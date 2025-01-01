package enteties

import (
	"time"
)

type Message struct {
	DeviceID        int
	Timestamp       time.Time
	SomeUsefulField int
}

func NewMessage(deviceID int, usefulField int) *Message {
	return &Message{
		DeviceID:        deviceID,
		Timestamp:       time.Now(),
		SomeUsefulField: usefulField,
	}
}
