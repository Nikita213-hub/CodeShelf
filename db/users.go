package db

import (
	"errors"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
)

func (db *Db) GetUser(username string) (*Models.User, error) {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", username)
	r := db.conn.QueryRow(db.ctx, queryString)
	var user Models.User
	err := r.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func (db *Db) AddUser(username string, password string) (*Models.User, error) {
	queryString := fmt.Sprintf("INSERT INTO users(username, password) VALUES('%s', '%s')", username, password)
	user, err := db.GetUser(username)
	if err == nil && user != nil {
		return nil, errors.New("user is already exists")
	}
	tag, err := db.conn.Exec(db.ctx, queryString)
	if err != nil {
		return nil, err
	}
	_ = tag
	return &Models.User{
		Username: username,
		Password: password,
	}, nil
}
