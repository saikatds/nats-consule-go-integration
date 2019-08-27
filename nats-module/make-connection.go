package nats_module

import (
	"github.com/nats-io/go-nats"
	"log"
)

func GetConnection() *nats.Conn {

	// connecting to nats server
	natsConnection  , error := nats.Connect(nats.DefaultURL)

	if error != nil {
		log.Fatal("Could not connect to nats server due to ",error)
	}else{
		log.Println("Successfully connected to nats server with URL ",nats.DefaultURL)
	}

	return natsConnection
}
