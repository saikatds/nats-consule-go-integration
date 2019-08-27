package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main(){

	//Getting a new client
	client , err := api.NewClient(api.DefaultConfig())

	if err != nil {
		panic(err)
	}

	agent :=  client.Agent()
	// Get a handle to the KV API
	kv := client.KV()

	p := &api.KVPair{Key: "REDIS_MAX-CLIENTS", Value: []byte("1000")}

	_, err = kv.Put(p, nil)

	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAX-CLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(" KV :", pair.Key, pair.Value)
}
