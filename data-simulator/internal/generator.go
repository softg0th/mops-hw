package internal

import (
	enteties "data-simulator/internal/enteties"
	"github.com/panjf2000/ants"
	"math/rand/v2"
	"sync"
	"time"
)

func StartGeneratingMessages(payload *enteties.Payload) {
	defer ants.Release()
	runTimes := 10000
	var wg sync.WaitGroup

	runDeviceTask := func() {
		defer wg.Done()
		deviceTask(payload.CountOfMessages, 1)
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(runDeviceTask)
	}
	wg.Wait()
}

func deviceTask(countOfMsg, deviceID int) {
	msg := enteties.NewMessage(deviceID, rand.IntN(100))
	sendMessage(msg)
	time.Sleep(time.Duration(countOfMsg))
}
