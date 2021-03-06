package psqlRepo

import (
	"context"
	"evc/entity/user"
	"evc/repository"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

const (
	SelectByEmail    = `SELECT id, first_name, last_name, email, date FROM user_table WHERE email=$1`
	SelectByDate     = `SELECT * FROM user_table WHERE date BETWEEN ? AND ? VALUES($1, $2)`
	SelectAll        = `SELECT * FROM user_table LIMIT 10 OFFSET = $1`
	SelectAllByLimit = `SELECT * FROM user_table LIMIT = $1`
	SelectAllByPage  = `SELECT * FROM user_table LIMIT 10 OFFSET = $1`
	InsertStmt       = `INSERT INTO user_table(id, first_name , last_name, email, date) VALUES($1, $2, $3, $4, $5)`
)

type psql struct {
	driver *pgx.Conn
}

func (p *psql) Save(user2 *user.User) (*user.User, error) {
	_, err := p.driver.Exec(context.Background(), InsertStmt, user2.Id, user2.FirstName, user2.LastName, user2.Email, user2.Date)
	if err != nil {
		return nil, err
	}
	return user2, nil
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
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date)
	if err != nil {
		log.Printf("Error fetch user: %v", err)
		return nil, err
	}
	return &user, err

}

func (p *psql) GetByDate(limit, offset string, start, end string) ([]user.User, error) {
	var users []user.User
	qbd := fmt.Sprintf("SELECT * FROM user_table WHERE date BETWEEN '%v' AND '%v' LIMIT %v OFFSET %v", start, end, limit, offset)
	fmt.Println(qbd)
	rows, err := p.driver.Query(context.Background(), qbd)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user user.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (p *psql) GetAllUsers(limit, offset string) ([]user.User, error) {
	var users []user.User
	queryBuilder := fmt.Sprintf("SELECT id, first_name,last_name,email,date FROM user_table LIMIT %v OFFSET %v", limit, offset)
	query, err := p.driver.Query(context.Background(), queryBuilder)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	for query.Next() {
		var user user.User
		err := query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date)
		if err != nil {
			log.Printf("Error fetching users: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewPsql(driver *pgx.Conn) repository.UserRepoInterface {
	return &psql{driver: driver}
}
