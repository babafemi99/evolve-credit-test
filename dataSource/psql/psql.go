package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var conn *pgx.Conn
var (
	dsn = os.Getenv("DSN")
)

func init() {
	var err error
	err = godotenv.Load("C:\\Users\\OREOLUWA\\Documents\\evolve-credit-test\\prod.env")
	if err != nil {
		log.Println(err)
	}
	//dsn := os.Getenv("DSN")
	conn, err = pgx.Connect(context.Background(), "postgres://webuser01:oooooooo@database-1.cv1yprcb749w.us-east-1.rds.amazonaws.com:5432/evcdb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	_, Execerr := conn.Exec(context.Background(), "create table if not exists user_table(id varchar(255), first_name varchar(255), last_name varchar(255), email varchar(255), date date)")
	if Execerr != nil {
		log.Fatalf("Error creating table in DB: %v", Execerr)
	}
	if err = conn.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected successfully")
}
func NewPql(dsn string) {
	var err error
	conn, err = pgx.Connect(context.Background(), "test string")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
}

func GetConn() *pgx.Conn {
	return conn
}
