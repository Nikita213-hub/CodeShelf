package db

import (
	"errors"
	"fmt"
)

func (db *Db) AddSession(sessionId int) error {
	queryStr := fmt.Sprintf("INSERT INTO sessions(id) VALUES('%d')", sessionId)
	tag, err := db.conn.Exec(db.ctx, queryStr)
	_ = tag
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) GetSession(sessionId int) error {
	queryStr := fmt.Sprintf("SELECT * FROM sessions WHERE id=%d", sessionId)
	tag, err := db.conn.Exec(db.ctx, queryStr)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("there is no session with such id")
	}
	return nil
}
