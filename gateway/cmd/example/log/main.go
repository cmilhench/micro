package main

import (
	"fmt"
	"log"

	"gateway/fabric"
)

func main() {
	forever := make(chan bool)

	err := fabric.ListenAndServe("fabric", "amqp://guest:guest@localhost:5672/", "log", "fabric.#", handle)
	if err != nil {
		panic(err)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func handle(req *fabric.Request) *fabric.Response {
	fmt.Printf(" [ ] %s([]byte{%d})\n", req.RoutingKey, len(req.Body))
	return nil
}
