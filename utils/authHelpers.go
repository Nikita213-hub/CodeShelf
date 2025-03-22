package utils

import (
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
	"net/http"
	"strconv"
)

func AuthCheck(r *http.Request, ua Models.IAuthStorageController) (isAuth bool, err error) {
	sessionId, err := r.Cookie("sessionId")
	if err != nil {
		return false, err
	}
	sIdInt, err := strconv.Atoi(sessionId.Value)
	fmt.Println(sIdInt)
	if err != nil {
		return false, err
	}
	err = ua.GetSession(sIdInt)
	if err != nil {
		return false, err
	}
	return true, nil
}
