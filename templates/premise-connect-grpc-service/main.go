package main

import (
	"connectrpc.com/connect"
	"context"
	greetv1 "example.com/premise-module/apps/premise-connect-grpc-service/gen/greet/v1"
	"example.com/premise-module/apps/premise-connect-grpc-service/gen/greet/v1/greetv1connect"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Private helpers ------------------------------------------------------------
type greeter struct{}

func (greeter) Greet(ctx context.Context, r *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	return connect.NewResponse(&greetv1.GreetResponse{Message: "Hello, " + r.Msg.Name + "!"}), nil
}
func handler() http.Handler {
	mux := http.NewServeMux()
	path, h := greetv1connect.NewGreetServiceHandler(greeter{})
	mux.Handle(path, h)
	return mux
}
func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	server := &http.Server{Addr: addr, Handler: handler(), ReadHeaderTimeout: 5 * time.Second, Protocols: new(http.Protocols)}
	server.Protocols.SetHTTP1(true)
	server.Protocols.SetUnencryptedHTTP2(true)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdown)
	}()
	log.Printf("listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
