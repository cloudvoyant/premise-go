package main

import (
	"connectrpc.com/connect"
	"context"
	greetv1 "example.com/premise-module/gen/greet/v1"
	"example.com/premise-module/gen/greet/v1/greetv1connect"
	"net/http/httptest"
	"testing"
)

// Public API -----------------------------------------------------------------
func TestProtocols(t *testing.T) {
	s := httptest.NewUnstartedServer(handler())
	s.EnableHTTP2 = true
	s.StartTLS()
	defer s.Close()
	for _, opts := range [][]connect.ClientOption{nil, {connect.WithGRPC()}, {connect.WithGRPCWeb()}} {
		c := greetv1connect.NewGreetServiceClient(s.Client(), s.URL, opts...)
		r, err := c.Greet(context.Background(), connect.NewRequest(&greetv1.GreetRequest{Name: "Ada"}))
		if err != nil {
			t.Fatal(err)
		}
		if r.Msg.Message != "Hello, Ada!" {
			t.Fatal(r.Msg.Message)
		}
	}
}
