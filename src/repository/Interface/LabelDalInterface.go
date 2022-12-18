package dal

import labelReq "backend/src/dto/request"

// dal -> data access layer - camada de acesso a dados
type LabelDataAccessLayerInterface interface {

	// CreateLabel cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateLabel(req labelReq.LabelRequestDto) (Label, error)

	ReadLabel(labelId string) (Label, error)

	UpdateLabel(labelID string, req UpdateLabelRequest) (Label, error)

	DeleteLabel(labelID string) error

	ListAllLabel(req ListLabelRequest) ([]Label, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}