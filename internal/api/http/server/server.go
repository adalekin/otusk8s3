package server

import (
	"fmt"
	"net/http"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/api/http/probes"
	"github.com/adalekin/otusk8s3/internal/api/http/users"
	"github.com/adalekin/otusk8s3/internal/appenv"
	"github.com/gorilla/mux"
)

func MakeServer(appEnv appenv.AppEnv) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/liveness", common.MakeHandler(appEnv, probes.LivenessHandler)).Methods("GET")
	r.HandleFunc("/users", common.MakeHandler(appEnv, users.CreateUserHandler)).Methods("POST")
	r.HandleFunc("/users", common.MakeHandler(appEnv, users.ListUsersHandler)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/", common.MakeHandler(appEnv, users.GetUserHandler)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/", common.MakeHandler(appEnv, users.UpdateUserHandler)).Methods("PATCH")
	r.HandleFunc("/users/{id:[0-9]+}/", common.MakeHandler(appEnv, users.DeleteUserHandler)).Methods("DELETE")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appEnv.Config.Port),
		Handler: r,
	}

	return srv
}
