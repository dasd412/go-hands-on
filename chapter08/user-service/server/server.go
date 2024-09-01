package main

import (
	"context"
	"errors"
	users "github.com/dasd412/user-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

type userService struct {
	users.UnimplementedUsersServer
}

func (s *userService) GetUser(
	ctx context.Context,
	in *users.UserGetRequest,
) (*users.UserGetReply, error) {
	log.Printf(
		"Received request for user with Email: %s Id: %s\n",
		in.Email,
		in.Id,
	)
	components := strings.Split(in.Email, "@")

	if len(components) != 2 {
		return nil, errors.New("invalid email")
	}

	u := users.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       36,
	}

	return &users.UserGetReply{User: &u}, nil
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":50051"
	}

	lis, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	users.RegisterUsersServer(s, &userService{})
	log.Fatal(s.Serve(lis))
}
