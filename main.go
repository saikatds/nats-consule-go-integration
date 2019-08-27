package main

import (
	"github.com/sdeysarkar/nats-consule-go-integration/nats-module"
	"github.com/sdeysarkar/nats-consule-go-integration/consul"
	"log"
)

func main(){

	//Registering Request Model
	requestService := new(consul.Service)
	requestService.Name="request-service"

	requestService, errRequest := consul.RegisterService(requestService)
	if errRequest != nil {
		log.Fatal(errRequest)
	}

	log.Println(" Request service Successfully Registered to Consul")

	//Registering Reply Model
	replyService := new(consul.Service)
	replyService.Name="reply-service"

	replyService, errReply := consul.RegisterService(replyService)
	if errReply != nil {
		log.Fatal(errReply)
	}

	log.Println("Reply service Successfully Registered to Consul")

	nats_module.ReplyModel()
	nats_module.RequestModel()


}
