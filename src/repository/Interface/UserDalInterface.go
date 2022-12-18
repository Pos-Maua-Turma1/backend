package dal

import (
	"backend/src/domain"
	"backend/src/dto"
)

// dal -> data access layer - camada de acesso a dados
type UserDataAccessLayerInterface interface {

	// CreateUser cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateUser(req dto.UserRequestDto) (domain.User, error)

	ReadUser(userId string) (domain.User, error)

	UpdateUser(userID string, req dto.UserRequestDto) (domain.User, error)

	DeleteUser(userID string) error

	// ListAllUser(req ListUserRequest) ([]domain.User, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}