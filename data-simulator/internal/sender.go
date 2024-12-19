package internal

import (
	enteties "data-simulator/internal/enteties"
	"log"
)

func sendMessage(message *enteties.Message) {
	log.Println(message)
}
