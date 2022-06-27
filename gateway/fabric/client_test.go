package fabric

import (
	"strconv"
	"testing"
)

func TestPublish(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	in := 12

	req := []byte(strconv.Itoa(in))
	res, err := Publish("fabric", "amqp://guest:guest@localhost:5672/", "fabric.fib.get", "text/plain", req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v, %+v\n", res.Request, res.Body)

	out, err := strconv.Atoi(string(res.Body))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v > %+v\n", in, out)
}
