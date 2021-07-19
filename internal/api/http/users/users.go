package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/appenv"
)

func ListUsersHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	status := common.Error{Code: "OK"}

	body, err := json.Marshal(status)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
