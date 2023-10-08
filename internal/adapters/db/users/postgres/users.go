package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var LoginAlreadyUsedErr = errors.New("login already used by another user")
var WrongLoginOrPasswordErr = errors.New("wrong login or password")

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type UserStorage struct {
	conn *pgx.Conn
	log  Logger
}

func NewUserStorage(conn *pgx.Conn, l Logger) *UserStorage {
	return &UserStorage{
		conn: conn,
		log:  l,
	}
}
