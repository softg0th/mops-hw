package enteties

type Payload struct {
	CountOfDevices   int
	MessageFrequency int
}

func NewPayload(devices int, messages int) *Payload {
	return &Payload{
		CountOfDevices:   devices,
		MessageFrequency: messages,
	}
}
