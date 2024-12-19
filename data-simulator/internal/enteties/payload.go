package enteties

type Payload struct {
	CountOfDevices  int
	CountOfMessages int
}

func NewPayload(devices int, messages int) *Payload {
	return &Payload{
		CountOfDevices:  devices,
		CountOfMessages: messages,
	}
}
