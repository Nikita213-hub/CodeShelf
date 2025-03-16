package db

import (
	"errors"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
)

func (db *Db) GetUser(username string) (*Models.User, error) {
	queryString := fmt.Sprintf("SELECT username FROM users WHERE username='%s'", username)
	tag, err := db.conn.Exec(db.ctx, queryString)
	fmt.Println(err)
	if err != nil || tag.RowsAffected() == 0 {
		return nil, err
	}
	fmt.Println(tag.RowsAffected())
	_ = tag
	return &Models.User{}, nil
}

func (db *Db) AddUser(username string) (*Models.User, error) {
	queryString := fmt.Sprintf("INSERT INTO users(username) VALUES('%s')", username)
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
	}, nil
}
