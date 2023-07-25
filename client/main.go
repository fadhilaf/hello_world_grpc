package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc_client/hello"
)

func main() {
	//baca intro.txt utk mahami lagi grpc

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// buat stub client, (yg abstraksi request, response, ke server)
	stub := hello.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //kalo timeout satu detik, rpc ny dicancel
	defer cancel()

	// panggil rpc Hello, menggunakan stub client
	r, err := stub.Hello(ctx, &hello.Void{}) //abstraksi ny, jadi seakan akan kito panggil fungsi dari dalam kode ini dewek
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting From Hello: %s", r.Message) //print response dari server

	// panggil rpc HelloAllTypes
	r2, err := stub.HelloAllTypes(ctx, &hello.HelloAllTypesRequest{
		Id:      1,
		Name:    "Hello",
		Salary:  1000.0,
		Active:  true,
		Strings: []string{"a", "b", "c"},
		Objects: []*hello.HelloObject{
			{
				Id:   23,
				Name: "Fadhil",
			},
			{
				Id:   24,
				Name: "Fadhil2",
			},
		}},
	)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting From HelloAllTypes: %s", r2.Message) //print response dari server
}
