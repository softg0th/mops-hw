package generator

import (
	enteties "data-simulator/internal/enteties"
	"data-simulator/internal/network"
	"github.com/panjf2000/ants"
	"math/rand/v2"
	"sync"
	"time"
)

func StartGeneratingMessages(payload *enteties.Payload, rpc *network.RPCConn) {
	defer ants.Release()
	var wg sync.WaitGroup

	runDeviceTask := func() {
		defer wg.Done()
		deviceTask(payload.CountOfMessages, 1, rpc)
	}

	for i := 0; i < payload.CountOfMessages; i++ {
		wg.Add(1)
		_ = ants.Submit(runDeviceTask)
	}
	wg.Wait()
}

func deviceTask(countOfMsg, deviceID int, rpc *network.RPCConn) {
	msg := enteties.NewMessage(deviceID, rand.IntN(100))
	rpc.StreamRequest(msg)
	time.Sleep(time.Duration(countOfMsg))
}
