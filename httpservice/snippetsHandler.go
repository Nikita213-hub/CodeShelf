package httpservice

import (
	"encoding/json"
	"fmt"
	"github.com/Nikita213-hub/CodeShelf/Models"
	"github.com/Nikita213-hub/CodeShelf/utils"
	"net/http"
	"strconv"
)

func newSnippet(sc Models.ISnippetsStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.FormValue("codesn")

		//fmt.Println(f)
		//file, handler, err := r.FormFile("form-id")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//_ = file
		//_ = handlerTimeline of the most recent commits to this repository and its network ordered by most recently pushed to.
		fp, err := utils.NewFile("go")
		if err != nil {
			fmt.Println(err)
		}
		_ = fp
		err = utils.WriteToFile(fp, payload)
		if err != nil {
			fmt.Println(err)
		}
		_, err = sc.NewSnippet(2, 2, "go", fp.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getSnippet(sc Models.ISnippetsStorageController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snippetId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid snippet ID", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		snippet, err := sc.GetSnippet(snippetId)
		if err != nil {
			http.Error(w, "Error occured while getting snippet", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		code, err := utils.GetFileContent(snippet.FileName)
		if err != nil {
			http.Error(w, "Error occured while getting snippet", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		resPayload := struct {
			Code string
			*Models.Snippet
		}{
			code,
			snippet,
		}
		fmt.Println(resPayload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(resPayload)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
	}
}
