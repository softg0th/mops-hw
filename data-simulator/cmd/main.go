package main

import (
	"bufio"
	exceptions "data-simulator/cmd/exceptions"
	"data-simulator/internal"
	"data-simulator/internal/enteties"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello world! Input devices count:")

	reader := bufio.NewReader(os.Stdin)
	countOfDevicesInput, _ := reader.ReadString('\n')
	countOfDevicesInput = strings.TrimSpace(countOfDevicesInput)
	countOfDevices, err := strconv.Atoi(countOfDevicesInput)
	if err != nil {
		handleError(&exceptions.CMDError{Field: "Device Count", Message: "invalid number format"})
		return
	}

	fmt.Println("Ok, Now input message frequency:")

	frequencyOfMessageInput, _ := reader.ReadString('\n')
	frequencyOfMessageInput = strings.TrimSpace(countOfDevicesInput)
	frequencyOfMessage, err := strconv.Atoi(frequencyOfMessageInput)

	if err != nil {
		handleError(&exceptions.CMDError{Field: "Message Frequency", Message: "invalid number format"})
		return
	}

	fmt.Println("Great! Now starting generating payload...")

	payload := enteties.NewPayload(countOfDevices, frequencyOfMessage)
	internal.StartGeneratingMessages(payload)
}

func handleError(err error) {
	fmt.Println(err)
	time.Sleep(time.Second)
	os.Exit(1)
}
