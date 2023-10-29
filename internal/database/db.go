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

//with any (working)
type Pool interface {
	*sql.DB | *mongo.Client | any
}

type DataPool[p Pool] struct {
	Pool p
}

type Database[p Pool] struct {
	DataType string
	Pool 	p
}

//working as long as any is present
func New[p Pool](dsn, dbType string) (*Database[Pool], error) {
	var (
		dbpool Pool
		err error
		database Database[Pool]
	)
	//find the type of db, then connect to it
	//maybe try to get rid of any inside pool, and do type asserition here instead
	switch strings.ToLower(dbType) {
	case "postgres", "postgresql":
		dbpool, err = NewPostgres(dsn)
		if err != nil {
			return nil, err
		}
		database.Pool = dbpool
		return &database, nil

	case "mongo":
		dbpool, err = NewMongo(dsn)
		if err != nil {
			return nil, err
		}
		database.Pool = dbpool
		return &database, nil
	default:
		return nil, errors.New("dbType is not supported")
	}
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

func NewMongo(dsn string) (*mongo.Client, error) {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

