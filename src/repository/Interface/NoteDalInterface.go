package dal

import (
	"backend/src/domain"
	"backend/src/dto"
)

// dal -> data access layer - camada de acesso a dados
type NoteDataAccessLayerInterface interface {

	// CreateNote cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateNote(req dto.NoteRequestDto) (domain.Note, error)

	ReadNote(noteId string) (domain.Note, error)

	UpdateNote(noteID string, req dto.NoteRequestDto) (domain.Note, error)

	DeleteNote(noteID string) error

	// ListAllNote(req ListNoteRequest) ([]Note, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}