package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

var conn *pgx.Conn

func init() {
	var err error
	conn, err = pgx.Connect(context.Background(), "test string")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

}

func GetConn() *pgx.Conn {
	return conn
}
