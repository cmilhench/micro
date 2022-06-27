package main

import (
	"consignments"
	"consignments/cmd/amqp/handler"
	"consignments/repository/memory"
	"gateway/fabric"
	"log"
)

// https://ewanvalentine.io/microservices-in-golang-part-1/

func main() {
	forever := make(chan bool)

	service := &consignments.Service{Repo: memory.New()}
	handle := handler.HandleRPC(service)
	err := fabric.ListenAndServe("fabric", "amqp://guest:guest@localhost:5672/", "consignments", "fabric.consignments.*", handle)
	if err != nil {
		panic(err)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
