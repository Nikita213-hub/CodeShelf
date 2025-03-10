package httpservice

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	sbc, err := res.Write([]byte("Hello\n"))
	if err != nil {
		fmt.Println("s", sbc)
	}
	return
}
