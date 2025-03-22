package httpservice

import (
	"encoding/json"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
	"net/http"
	"strconv"
	"time"
)

type AuthUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(ua Models.IAuthStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usr AuthUserReq
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			fmt.Println(err)
		}
		newUser, err := ua.AddUser(usr.Username, usr.Password)
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

func SignIn(ua Models.IAuthStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData AuthUserReq
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			fmt.Println(err)
		}
		user, err := ua.GetUser(userData.Username)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			return
		}
		if user.Password != userData.Password {
			w.WriteHeader(403)
			return
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
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println(err)
		}
	}
}
