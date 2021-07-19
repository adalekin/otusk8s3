package common

import (
	"net/http"

	"github.com/adalekin/otusk8s3/internal/appenv"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, appenv.AppEnv)

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(appEnv appenv.AppEnv, fn func(http.ResponseWriter, *http.Request, appenv.AppEnv)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, appEnv)
	}
}
