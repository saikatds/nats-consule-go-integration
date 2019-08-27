package nats_module

import (
	"github.com/nats-io/go-nats"
	"log"
)

func ReplyModel() {
	// connecting to nats server
	natConnection := GetConnection()

	//declaring subject and reply
	sub := "saikat"

	//Subscribing to the subject
	natConnection.Subscribe(sub , func(msg *nats.Msg) {
		log.Println("Message received  for subject ", sub , " is : " ,string(msg.Data ))
		//giving reply
		natConnection.Publish(msg.Reply , []byte (string(msg.Data)+" EFG") )
	})

	log.Println("Listening to subject ", sub)

	// Keep the connection alive
	//runtime.Goexit()
}
