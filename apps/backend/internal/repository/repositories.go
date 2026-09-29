package repository

import "github.com/PrathmeshAdhav2006/go-boilerplate/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
