package rpc

import (
	"fmt"
	"testing"
	"time"
)

func TestTrasport_StartStop(t *testing.T) {
	trasport := NewTCPTransport("127.0.0.1:0", time.Second, nil)
	go func() {
		// this is the server side
		// it should read message from the consumer channel and reply to them.
		rpc := <-trasport.Consumer()
		fmt.Printf("consumer got rpc command %#v\n", rpc.Command)
		rpc.Respond(&EchoResponse{"FOO"}, nil)
	}()

	// this is the client side
	res, err := trasport.Echo(trasport.LocalAddr(), "foo")
	if err != nil {
		t.Errorf("error %#v\n", err)
	}
	if res != "FOO" {
		t.Errorf("expected FOO instead %v", res)
	}
	fmt.Printf("returns %v\n", res)
}