package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(a *Account) error
	GetAccountById(id int) (*Account, error)
	UpdateAccount(a *Account) error
	DeleteAccount(id int) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	constr := "user=postgres dbname=postgres password=qwertyuiop sslmode=disable"

	db, err := sql.Open("postgres", constr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
func (s *PostgresStorage) CreateAccount(a *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStorage) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}
func (s *PostgresStorage) createAccountTable() error {
	query := "create table if not exists account (id serial primary key, first_name varchar(50), last_name varchar(50), number serial, balance serial, created_at timestamp )"
	_, err := s.db.Exec(query)
	return err

}
