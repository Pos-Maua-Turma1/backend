package dal

// dal -> data access layer - camada de acesso a dados
type UserDataAccessLayerInterface interface {

	// CreateUser cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateUser(req CreateUserRequest) (User, error)

	ReadUser(userId string) (User, error)

	UpdateUser(userID string, req UpdateUserRequest) (User, error)

	DeleteUser(userID string) error

	ListAllUser(req ListUserRequest) ([]User, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}