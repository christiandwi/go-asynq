package guest

import (
	"github.com/christiandwi/go-asynq/database"
	"github.com/hibiken/asynq"
)

type service struct {
	db    *database.Database
	asynq *asynq.Server
}

func NewGuestService(db *database.Database, asynq *asynq.Server) Service {
	return &service{
		db:    db,
		asynq: asynq,
	}
}
