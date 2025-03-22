package db

import (
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
)

func (db *Db) NewSnippet(ownerId, pLangId int, password, fileName string) (*Models.Snippet, error) {
	queryString := fmt.Sprintf("INSERT INTO snippets(owner_id, password, file, p_lang) VALUES('%d', '%s', '%s', '%d') RETURNING id",
		ownerId, password, fileName, pLangId)
	row := db.conn.QueryRow(db.ctx, queryString)
	var queryRes int
	err := row.Scan(&queryRes)
	if err != nil {
		return nil, err
	}
	return &Models.Snippet{
		Id:       queryRes,
		OwnerId:  ownerId,
		FileName: fileName,
		Password: password,
		PLang:    pLangId,
	}, nil
}

func (db *Db) UploadSnippet(snippetId int, snippetCode string) error { return nil }

func (db *Db) GetSnippet(snippetId int) (*Models.Snippet, error) {
	queryString := fmt.Sprintf("SELECT * FROM snippets WHERE id='%d'", snippetId)
	row := db.conn.QueryRow(db.ctx, queryString)
	var queryRes Models.Snippet
	err := row.Scan(&queryRes.Id, &queryRes.OwnerId, &queryRes.Password, &queryRes.FileName, &queryRes.PLang)
	if err != nil {
		return nil, err
	}
	return &queryRes, nil
}
