package Models

type IAuthStorageController interface {
	AddUser(username string, password string) (*User, error)
	GetUser(username string) (*User, error)
	AddSession(sessionId int) error
	GetSession(sessionId int) error
}

type ISnippetsStorageController interface {
	NewSnippet(ownerId, pLangId int, password, fileName string) (*Snippet, error)
	UploadSnippet(snippetId int, snippetCode string) error
	GetSnippet(snippetId int) (*Snippet, error)
}
