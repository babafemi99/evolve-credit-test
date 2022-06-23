package psqlRepo

import (
	"context"
	"evc/entity/user"
	"evc/repository"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

const (
	SelectByEmail    = `SELECT * FROM user_table WHERE email IS = $1`
	SelectByDate     = `SELECT * FROM user_table WHERE date BETWEEN ? AND ? VALUES($1, $2)`
	SelectAll        = `SELECT * FROM user_table LIMIT 10 OFFSET = $1`
	SelectAllByLimit = `SELECT * FROM user_table LIMIT = $1`
	SelectAllByPage  = `SELECT * FROM user_table LIMIT 10 OFFSET = $1`
)

type psql struct {
	driver *pgx.Conn
}

func (p *psql) GetAllUsersByLimit(limit int) (*user.Users, error) {
	var users user.Users
	query, err := p.driver.Query(context.Background(), SelectAllByLimit, limit)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	for query.Next() {
		var user user.User
		err := query.Scan(&user)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users.Add(&user)
	}
	return &users, nil
}

func (p *psql) GetAllUsersByPage(page int) (*user.Users, error) {
	var users user.Users
	query, err := p.driver.Query(context.Background(), SelectAllByPage, page)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	for query.Next() {
		var user user.User
		err := query.Scan(&user)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users.Add(&user)
	}
	return &users, nil
}

func (p *psql) GetByEmail(email string) (*user.User, error) {
	var user user.User
	row := p.driver.QueryRow(context.Background(), SelectByEmail, email)
	err := row.Scan(&user)
	if err != nil {
		log.Printf("Error fetch user: %v", err)
		return nil, err
	}
	return &user, err

}

func (p *psql) GetByDate(limit, offset string, start, end time.Time) (*user.Users, error) {
	var users user.Users
	queryBuilder := "SELECT * FROM user_table WHERE date BETWEEN ? AND ? VALUES($1, $2) LIMIT" + limit + "OFFSET" + offset
	query, err := p.driver.Query(context.Background(), queryBuilder, start, end)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	for query.Next() {
		var user user.User
		err := query.Scan(&user)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users.Add(&user)
	}
	return &users, nil
}

func (p *psql) GetAllUsers(limit, offset string) (*user.Users, error) {
	var users user.Users
	queryBuilder := "SELECT * FROM user_table LIMIT" + limit + "OFFSET" + offset
	query, err := p.driver.Query(context.Background(), queryBuilder)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	for query.Next() {
		var user user.User
		err := query.Scan(&user)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users.Add(&user)
	}
	return &users, nil
}

func NewPsql(driver *pgx.Conn) repository.UserRepoInterface {
	return &psql{driver: driver}
}
