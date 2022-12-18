package dal

import (
	"backend/src/domain"
	"backend/src/dto"
)
// dal -> data access layer - camada de acesso a dados
type LabelDataAccessLayerInterface interface {

	// CreateLabel cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateLabel(req dto.LabelRequestDto) (domain.Label, error)

	ReadLabel(labelId string) (domain.Label, error)

	UpdateLabel(labelID string, req dto.LabelRequestDto) (domain.Label, error)

	DeleteLabel(labelID string) error

	// ListAllLabel(req ListLabelRequest) ([]domain.Label, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}