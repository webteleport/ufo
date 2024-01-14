package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "github.com/webteleport/ufo/apps/buf/gen/greet/v1"
	"github.com/webteleport/ufo/apps/buf/gen/greet/v1/greetv1connect"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	/*
		err := http.ListenAndServe(
			"localhost:8080",
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
	*/
	log.Println(`curl https://buf.ufo.k0s.io/greet.v1.GreetService/Greet --data '{"name": "Jane"}' --header "Content-Type: application/json"`)
	err := wtf.Serve(
		"https://ufo.k0s.io/buf?clobber=buf",
		// Use h2c so we can serve HTTP/2 without TLS.
		utils.GinLoggerMiddleware(utils.AllowAllCorsMiddleware(h2c.NewHandler(mux, &http2.Server{}))),
	)
	if err != nil {
		log.Fatalf("listen failed: %v\n", err)
	}
}
