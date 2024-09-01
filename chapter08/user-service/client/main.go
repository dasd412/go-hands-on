package main

import (
	"context"
	"fmt"
	users "github.com/dasd412/user-service/service"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(
			"Must specify a gRPC server address",
		)
	}
	addr := os.Args[1]

	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := users.NewUsersClient(conn)

	result, err := c.GetUser(context.Background(), &users.UserGetRequest{Email: "jane@doe.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(
		os.Stdout, "User: %s %s\n",
		result.User.FirstName,
		result.User.LastName,
	)
}
