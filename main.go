package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/twitchtv/twirp"

	pb "github.com/bkono/twirp-greeter/proto"
)

var (
	serverCommand = flag.NewFlagSet("server", flag.ExitOnError)
	port          = serverCommand.String("port", "", "port to listen on")

	clientCommand = flag.NewFlagSet("client", flag.ExitOnError)
	serverAddr    = clientCommand.String("addr", "", "addr of the service")
	name          = clientCommand.String("name", "John Doe", "name to greet")
)

var usage = `
usage: greeter <command> [<args>]
Command should be one of 'server' or 'client'.

Server requires the -port argument
Client requires the -addr argument
`

type server struct{}
type authIDKey struct{}

func withAuthID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("X-User-Auth")
		ctx := context.WithValue(req.Context(), authIDKey{}, auth)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func authIDFromCtx(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(authIDKey{}).(string)
	return id, ok
}

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	if authID, ok := authIDFromCtx(ctx); ok {
		log.Printf("[server] Greet received with authID: %s\n", authID)
	}
	if req.Name == "" {
		return nil, twirp.RequiredArgumentError("name")
	}

	return &pb.GreetResponse{Greeting: "Hello, " + req.Name + "!"}, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	switch os.Args[1] {
	case "server":
		serverCommand.Parse(os.Args[2:])
		runServer()
	case "client":
		clientCommand.Parse(os.Args[2:])
		doClientReq()
	default:
		fmt.Println(usage)
		os.Exit(2)
	}
}

func runServer() {
	if *port == "" {
		*port = os.Getenv("PORT")
		if *port == "" {
			fmt.Println("port is required via flag or environment variable for the server command")
			return
		}
	}
	// do the server stuff
	s := &server{}
	twirpHandler := pb.NewGreeterServer(s, nil)

	addr := ":" + *port
	log.Printf("[server] listening on addr: %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, withAuthID(twirpHandler)))
}

func doClientReq() {
	if *serverAddr == "" {
		fmt.Println("addr is required for client mode")
		return
	}

	// do the client stuff
	header := make(http.Header)
	header.Set("X-User-Auth", "1010101101001201")

	client := pb.NewGreeterProtobufClient(*serverAddr, &http.Client{})
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		log.Printf("[client] failed setting headers: %v\n", err)
		return
	}

	resp, err := client.Greet(ctx, &pb.GreetRequest{Name: *name})
	if err != nil {
		if twerr, ok := err.(twirp.Error); ok {
			switch twerr.Code() {
			case twirp.InvalidArgument:
				log.Printf("[client] woops, didn't provide a valid argument: %v\n", twerr.Msg())
			default:
				log.Printf("[client] error during Greet rpc: %v\n", twerr.Error())
			}
		}
	}

	log.Printf("[client] resp received: %v\n", resp)
}
