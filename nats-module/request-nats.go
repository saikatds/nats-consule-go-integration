package nats_module

import (
	"log"
	"time"
)

func RequestModel() {
	// connecting to nats server
	natConnection := GetConnection()

	//closing the connection after getting the reply
	defer natConnection.Close()

	sub , request := "saikat" , " 1123456 "

	//Requesting to the nats server and getting the reply
	reply , err := natConnection.Request(sub , []byte (request) , 100*time.Millisecond)

	//Checking if any error coming
	if err != nil {
		log.Fatal("Req-Reply model broke down due to ",err)
	}

	// Request and Reply data
	log.Println(" Request : ", string(request))
	log.Println("Reply : ", string(reply.Data))
}
