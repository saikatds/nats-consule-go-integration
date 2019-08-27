package main

import (
	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/gnatsd/test"
	"github.com/nats-io/go-nats"
	"github.com/sdeysarkar/nats-consule-go-integration/consul"
	"log"
	"runtime"
)

// RunDefaultServer will run a server on the default port.
func RunDefaultServer() *server.Server {
	return RunServerOnPort(nats.DefaultPort)
}

// RunServerOnPort will run a server on the given port.
func RunServerOnPort(port int) *server.Server {
	opts := test.DefaultTestOptions
	opts.Port = port
	return RunServerWithOptions(opts)
}

// RunServerWithOptions will run a server with the given options.
func RunServerWithOptions(opts server.Options) *server.Server {
	return test.RunServer(&opts)
}

func main(){

	//Registering Reply Model
	natsService := new(consul.Service)
	natsService.Name="nats-service"

	natsService, errReply := consul.RegisterService(natsService)
	if errReply != nil {
		log.Fatal(errReply)
	}

	log.Println("nats service Successfully Registered to Consul")

	server:= RunDefaultServer()

	isStarted := server.Start

	log.Println("Server status : ",isStarted)

	//keep the server alive
	runtime.Goexit()

}
