package enteties

import (
	"time"
)

type Message struct {
	deviceID        int
	timestamp       time.Time
	someUsefulField int
}

func NewMessage(deviceID int, usefulField int) *Message {
	return &Message{
		deviceID:        deviceID,
		timestamp:       time.Now(),
		someUsefulField: usefulField,
	}
}
