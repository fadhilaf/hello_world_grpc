package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"grpc_server/hello"
)

func main() {
	//baca intro.txt utk mahami lagi grpc

	// buat listener protokol tcp (berisi informasi2 utk server grpc)
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	// buat server grpc
	grpcServer := grpc.NewServer()

	// assign service ke server grpc
	hello.RegisterHelloServiceServer(grpcServer, &HelloService{})

	fmt.Println("Server running on port 9090")
	// jalankan server grpc
	if err := grpcServer.Serve(lis); err != nil {
		grpcServer.Stop()

		panic(err)
	}
}

// +++++++++++++ Bagian Definisi Service (function go, yg di assign ke grpc biar bisa di call dari client grpc ) +++++++++++++

// buat struct yg mengimplementasikan interface HelloServiceServer (berisi method2 service yg didefinisikan di proto)
type HelloService struct {
	hello.UnimplementedHelloServiceServer // unimplemented berisi method2 yg belum diimplementasikan (template ngebuat struct ny emang harus ada ini)
}

var _ hello.HelloServiceServer = &HelloService{}

// type request response dk perlu validasi, dan usdah didefinisikan di proto
func (s *HelloService) Hello(ctx context.Context, void *hello.Void) (*hello.HelloMessage, error) {
	message := "Hello, World! from service"

	return &hello.HelloMessage{Message: message}, nil
}

func (s *HelloService) HelloAllTypes(ctx context.Context, req *hello.HelloAllTypesRequest) (*hello.HelloAllTypesResponse, error) {
	id := strconv.Itoa(int(req.Id)) + " | "
	name := req.Name + " | "
	salary := strconv.FormatFloat(float64(req.Salary), 'f', -1, 64) + " | "
	active := strconv.FormatBool(req.Active) + " | "
	joinedStrings := strings.Join(req.Strings, ", ") + " | "

	var objectString []string
	for _, v := range req.Objects {
		objectId := strconv.Itoa(int(v.Id)) + "-"
		objectName := v.Name

		objectString = append(objectString, "("+objectId+objectName+")")
	}

	joinedObjects := strings.Join(objectString, ", ")

	message := id + name + salary + active + joinedStrings + joinedObjects

	return &hello.HelloAllTypesResponse{Message: message}, nil
}
