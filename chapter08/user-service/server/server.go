package server

import (
	"context"
	"github.com/dasd412/user-service/service"
)

type userService struct {
	service.UnimplementedUsersServer
}

func (s *userService) GetUser(
	ctx context.Context,

)
