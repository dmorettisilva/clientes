package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//DBConfig - contem os parametros para conexão do banco de dados oracle
type DBConfig struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbSid  string
}

//Config - struct para guardar configurações do aplicativo
type Config struct {
	DBConfig           *DBConfig
	BaseUrl            string
	Authorization      string
	AuthorizationAlpha string
}

func New() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "não pode carregar as váriaveis de conexão")
	}

	return &Config{
		DBConfig: &DBConfig{
			DbUser: os.Getenv("DB_USER"),
			DbPass: os.Getenv("DB_PASS"),
			DbHost: os.Getenv("DB_HOST"),
			DbPort: os.Getenv("DB_PORT"),
			DbSid:  os.Getenv("DB_SID"),
		},
		BaseUrl:       "https://localhost:8089",
		Authorization: "123456",
	}, nil
}

func (d *DBConfig) ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("godror", fmt.Sprintf("%s/%s@%s:%s/%s", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbSid))
	if err != nil {
		return nil, errors.Wrap(err, "erro ao abrir conexao")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "erro ao pingar no banco")
	}

	db.Mapper = reflectx.NewMapperTagFunc("db",
		nil,
		func(s string) string {
			return strings.ToUpper(s)
		},
	)

	sqlx.BindDriver("godror", sqlx.NAMED)

	return db, nil
}
