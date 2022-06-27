package main

import (
	"fmt"
	"log"
	"strconv"

	"gateway/fabric"
)

func main() {
	forever := make(chan bool)

	err := fabric.ListenAndServe("fabric", "amqp://guest:guest@localhost:5672/", "add", "fabric.add.one", handle)
	if err != nil {
		panic(err)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func handle(req *fabric.Request) *fabric.Response {
	in, err := strconv.Atoi(string(req.Body))
	if err != nil {
		res := &fabric.Response{Header: nil, Body: []byte(err.Error()), Request: req}
		return res
	}
	fmt.Printf(" [.] add(%d)\n", in)
	out := in + 1
	fmt.Printf(" [<] add(%d) %d\n", in, out)

	resp := []byte(strconv.Itoa(out))
	res := &fabric.Response{Header: nil, Body: resp, Request: req}
	return res
}
