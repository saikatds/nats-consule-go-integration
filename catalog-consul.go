package main

import (
	"fmt"
	"github.com/drnic/consul-discovery"
)

func main() {
	client, _ := consuldiscovery.NewClient(consuldiscovery.DefaultConfig())

	services, _ := client.CatalogServices()
	//fmt.Printf("Services: %#v\n\n", services)

	for _, service := range services {
		serviceNodes, _ := client.CatalogServiceByName(service.Name)
		fmt.Println(service.Name, serviceNodes)
	}
}
