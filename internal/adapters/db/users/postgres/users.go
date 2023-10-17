package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var ErrLoginAlreadyUsed = errors.New("login already used by another user")
var ErrWrongLoginOrPassword = errors.New("wrong login or password")
var ErrUserDoseNotExist = errors.New("user dose not exist")
var ErrEmptyHistory = errors.New("empty history")

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
