package httpservice

import (
	"encoding/json"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
	"net/http"
	"strconv"
	"time"
)

type IAuthStorageController interface {
	AddUser(username string) (*Models.User, error)
	AddSession(sessionId int) error
	GetSession(sessionId int) error
}

type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(ua IAuthStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usr SignUpReq
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			fmt.Println(err)
		}
		newUser, err := ua.AddUser(usr.Username)
		if err != nil {
			w.WriteHeader(401)
			fmt.Println(err)
		}
		sessionId := int(time.Now().Unix())
		err = ua.AddSession(sessionId)
		if err != nil {
			fmt.Println(err)
		}
		cookie := http.Cookie{
			Name:     "sessionId",
			Value:    strconv.Itoa(sessionId),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
		err = json.NewEncoder(w).Encode(&newUser)
		if err != nil {
			fmt.Println(err)
		}
	}
}
