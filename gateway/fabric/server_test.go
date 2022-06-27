package fabric

// func Test(t *testing.T) {
// 	err := ListenAndServe("fabric", "amqp://guest:guest@localhost:5672/", "fibonacci", "fabric.fib.*", handle)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	time.Sleep(time.Millisecond * 10003)
// }

// func handle(req *Request) (*Response, error) {
// 	in, err := strconv.Atoi(string(req.Body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf(" [.] fib(%d)\n", in)
// 	out := fib.Fibonacci(in)
// 	fmt.Printf(" [<] fib(%d) %d\n", in, out)

// 	resp := []byte(strconv.Itoa(out))
// 	res := &Response{Header: nil, Body: resp, Request: req}
// 	return res, nil
// }
