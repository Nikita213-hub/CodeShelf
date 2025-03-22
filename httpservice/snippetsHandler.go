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
		payloadCode := r.FormValue("codesn")
		password := r.FormValue("password")
		pLangId := r.FormValue("prog_lang_id")
		pLangIdInt, err := strconv.Atoi(pLangId)
		if err != nil {
			http.Error(w, "Invalid programming language ID", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		userId := r.Context().Value("userId")
		var castedUserId int
		v, ok := userId.(int)
		if ok {
			castedUserId = v
		} else {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		//fmt.Println(f)
		//file, handler, err := r.FormFile("form-id")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//_ = file
		//_ = handler
		fp, err := utils.NewFile("go")
		if err != nil {
			http.Error(w, "Internal error while creating new file", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		_ = fp
		err = utils.WriteToFile(fp, payloadCode)
		if err != nil {
			http.Error(w, "Internal error while updating file", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		_, err = sc.NewSnippet(castedUserId, pLangIdInt, password, fp.Name())
		if err != nil {
			http.Error(w, "Internal while creating new record in database", http.StatusBadRequest)
			fmt.Println(err)
			return
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
