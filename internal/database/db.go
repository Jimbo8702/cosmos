package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pool interface {
	*sql.DB | *mongo.Client | any
}

type Database[p Pool] struct {
	DataType string
	Pool 	p
}

func New[p Pool](dsn, dbType string) (*Database[Pool], error) {
	var (
		dbpool Pool
		err error
		database Database[Pool]
	)
	
	//find the type of db, then connect to it
	switch strings.ToLower(dbType) {
	case "postgres", "postgresql":
		dbpool, err = NewPostgres(dsn)
		if err != nil {
			return nil, err
		}
		database.Pool = dbpool
		return &database, nil

	case "mongo":
		dbpool, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
		if err != nil {
			return nil, err
		}
		database.Pool = dbpool
		return &database, nil
	}
	return nil, errors.New("error loading db, could not find dbtype")
}

func NewPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
