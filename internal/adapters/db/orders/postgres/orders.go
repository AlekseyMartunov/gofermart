package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrOrderAlreadyCreated = errors.New("the order was already loaded")
var ErrEmptyResult = errors.New("empty result")
var ErrNotEnoughMoney = errors.New("not enough money")

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type OrderStorage struct {
	conn *pgxpool.Pool
	log  Logger
}

func NewOrderStorage(conn *pgxpool.Pool, l Logger) *OrderStorage {
	return &OrderStorage{
		conn: conn,
		log:  l,
	}
}
