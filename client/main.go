package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hello/proto"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// address     = "localhost:50051"
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	fmt.Println("vim-go-client")
	fmt.Println(address)

	tmpl := template.Must(template.ParseFiles("./index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Printf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			fmt.Printf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", res.GetMessage())
		tmpl.Execute(w, tmpl)

	})

	http.ListenAndServe(":8080", nil)

}
