package main

import (
	"context"
	"fmt"
	"log"
	"s0609-23/internal/authapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := authapi.NewAuthServiceClient(conn)

	response, err := client.UserLogin(context.Background(), &authapi.LoginRequest{
		Username: "alex",
		Password: "qwerty",
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}
}
