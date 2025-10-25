package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/viacheslav-korobeynikov/rally-app/config"
)

// Функция создания подключения к БД
func CreateDbPool(config *config.DbConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("Не удалось подключиться к БД")
		panic(err)
	}
	logger.Info().Msg("Подключились к БД")
	return dbpool
}
