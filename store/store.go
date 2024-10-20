package store

import (
	"TestTask/store/adapters/pg"
	"TestTask/store/usersS"
)

// Store - структура регистра, содержащая подключение и адаптер
type Store struct {
	Adapter *pg.PostgresAdapter
	Users   *usersS.Storage
}

// хранит экземпляр структуры Store
var Default *Store

func InitStore(cfg *Config) (*Store, error) {
	// Создаем адаптер для работы с базой данных PostgreSQL
	adapter, err := pg.NewAdapter(cfg.Postgres)
	if err != nil {
		return nil, err
	}
	users := usersS.NewStore(adapter.DB)
	if err != nil {
		return nil, err
	}
	// Возвращаем Store, содержащий адаптер
	return &Store{
		Adapter: adapter,
		Users:   users,
	}, nil
}
