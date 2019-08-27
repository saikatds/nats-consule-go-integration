package consul

import (
	consul "github.com/hashicorp/consul/api"
)

//Service model struct
type Service struct{
	Name string
	ConsulAgent *consul.Agent
}

func RegisterService (service *Service)(*Service, error){

	//getting the client
	client, err := consul.NewClient(consul.DefaultConfig())

	//Checking if any error comes or not
	if err != nil {
		return nil, err
	}

	//Assigning the agent to service struct
	service.ConsulAgent = client.Agent()

	//registering the service to consul
	serviceDef := &consul.AgentServiceRegistration{
		Name: service.Name ,
		Port:8085,
		Address:"http://localhost",
		//Address:
	}

	if  err := service.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}
	return service, nil
}


