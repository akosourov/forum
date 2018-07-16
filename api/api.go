package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akosourov/forum/data"
)

// GET /api/v1/threads
func Threads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	threads := data.ThreadList()
	data, err := json.Marshal(&threads)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, string(data))
}
