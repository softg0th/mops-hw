package main

import (
	"bufio"
	"data-simulator/internal/enteties"
	"data-simulator/internal/exceptions"
	"data-simulator/internal/generator"
	"data-simulator/internal/network"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Initializing connection...")
	rpc := network.NewRPCConn()

	fmt.Println("Hello world! Input devices count:")

	reader := bufio.NewReader(os.Stdin)
	countOfDevicesInput, _ := reader.ReadString('\n')
	countOfDevicesInput = strings.TrimSpace(countOfDevicesInput)
	countOfDevices, err := strconv.Atoi(countOfDevicesInput)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Device Count", Message: "invalid number format"})
		return
	}

	fmt.Println("Ok, Now input message frequency:")

	frequencyOfMessageInput, _ := reader.ReadString('\n')
	frequencyOfMessageInput = strings.TrimSpace(countOfDevicesInput)
	frequencyOfMessage, err := strconv.Atoi(frequencyOfMessageInput)

	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Message Frequency", Message: "invalid number format"})
		return
	}

	fmt.Println("Great! Now starting generating payload...")

	payload := enteties.NewPayload(countOfDevices, frequencyOfMessage)
	generator.StartGeneratingMessages(payload, rpc)
}
