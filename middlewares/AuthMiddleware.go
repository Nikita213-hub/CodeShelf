package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nikita213-hub/CodeShelf/Models"
	"github.com/Nikita213-hub/CodeShelf/utils"
)

func AuthMiddleware(handler http.HandlerFunc, ua Models.IAuthStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuth, err := utils.AuthCheck(r, ua)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			_, err := w.Write([]byte("Unauthorized error\n"))
			if err != nil {
				return
			}
			return
		}
		if !isAuth {
			w.WriteHeader(403)
			_, err := w.Write([]byte("Unauthorized error\n"))
			if err != nil {
				return
			}
			return
		}
		usrIdFrmCookie, err := r.Cookie("userId")
		if err != nil {
			w.WriteHeader(403)
			_, err := w.Write([]byte("Unauthorized error\n"))
			if err != nil {
				return
			}
			return
		}
		userId, err := strconv.Atoi(usrIdFrmCookie.Value)
		if err != nil {
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)
		rWithCtx := r.WithContext(ctx)
		handler.ServeHTTP(w, rWithCtx)
	}
}
