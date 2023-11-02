package storage

import (
	"database/sql"
	"fmt"

	"github.com/math-88/medic-online/internal/config"
	controller "github.com/math-88/medic-online/internal/controller/grpc"
	"github.com/math-88/medic-online/internal/usecase"
	"github.com/math-88/medic-online/storage/in_memory"
	"github.com/math-88/medic-online/storage/sqlc"
)

type Storage interface {
	controller.UserStore
	usecase.Repo
}

var _ Storage = &sqlc.Storage{}
var _ Storage = &in_memory.InMemoryStore{}

func New(cfg *config.Config) Storage {

	// InMemoryStore
	if cfg.App.IsDebug != nil && *cfg.App.IsDebug {
		return in_memory.NewInMemoryStore()
	}

	// mysql
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Database),
	)
	if err != nil {
		panic(err)
	}

	q := sqlc.New(db)
	wrapper := sqlc.NewWrapper(q)

	return wrapper
}
