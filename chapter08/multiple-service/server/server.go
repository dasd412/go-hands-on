package main

import (
	"context"
	"errors"
	"github.com/dasd412/multiple-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

type userService struct {
	service.UnimplementedUsersServer
}
type repoService struct {
	service.UnimplementedRepoServer
}

func (s *userService) GetUser(ctx context.Context, in *service.UserGetRequest) (*service.UserGetReply, error) {
	log.Printf("Received request for user with Email: %s Id: %s\n", in.Email, in.Id)

	components := strings.Split(in.Email, "@")

	if len(components) != 2 {
		return nil, errors.New("invalid email")
	}

	u := service.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       36,
	}
	return &service.UserGetReply{User: &u}, nil
}

func (s *repoService) GetRepos(ctx context.Context, in *service.RepoGetRequest) (*service.RepoGetReply, error) {
	log.Printf(
		"Received request for repo with CreateId: %s Id: %s\n",
		in.CreatorId,
		in.Id,
	)
	repo := service.Repository{
		Id:    in.Id,
		Name:  "test repo",
		Url:   "https://git.example.com/test/repo",
		Owner: &service.User{Id: in.CreatorId, FirstName: "Jane"},
	}
	r := service.RepoGetReply{Repo: []*service.Repository{&repo}}
	return &r, nil
}
func registerServices(s *grpc.Server) {
	service.RegisterUsersServer(s, &userService{})
	service.RegisterRepoServer(s, &repoService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
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
	registerServices(s)
	log.Fatal(startServer(s, lis))
}
