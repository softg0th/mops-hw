package core

type ObservedDevices struct {
	devices       map[int]int
	targetValue   int
	durationValue int
}

func NewObservedDevices(targetValue int, durationValue int) *ObservedDevices {
	return &ObservedDevices{
		devices:       make(map[int]int),
		targetValue:   targetValue,
		durationValue: durationValue,
	}
}

func (od *ObservedDevices) ProcessMessage(deviceID int, newValue int) (bool, bool) {
	if _, exists := od.devices[deviceID]; !exists {
		od.devices[deviceID] = 0
	}

	isMatchInstant := false
	isMatchWithDuration := false

	if newValue > od.targetValue {
		isMatchInstant = true
		od.devices[deviceID] += 1

		if od.devices[deviceID] >= od.durationValue {
			isMatchWithDuration = true
			od.devices[deviceID] = 0
		}
	} else {
		od.devices[deviceID] = 0
	}

	return isMatchInstant, isMatchWithDuration
}
