package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/appenv"
	"github.com/adalekin/otusk8s3/internal/service/usersvc"
)

type UserResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.WithError(err).Error("Bad request")
		return
	}

	createRequest := usersvc.CreateRequest{}

	err = json.Unmarshal(body, &createRequest)

	if err != nil {
		log.WithError(err).Error("Bad request")
		return
	}

	userService := usersvc.New(appEnv.RepoRegistry)
	user, err := userService.Create(req.Context(), createRequest)

	if err != nil {
		log.WithError(err).Error("Create user failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	userResponse := UserResponse(*user)

	_ = appEnv.Render.JSON(w, http.StatusCreated, userResponse)
}

func ListUsersHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	userService := usersvc.New(appEnv.RepoRegistry)
	users, err := userService.FindAll(req.Context())

	if err != nil {
		log.WithError(err).Error("List users failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	usersResponse := make([]interface{}, len(users))
	for i := range users {
		usersResponse[i] = UserResponse(*users[i])
	}

	_ = appEnv.Render.JSON(w, http.StatusOK, usersResponse)
}

func GetUserHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	findByIdRequest := usersvc.FindByIdRequest{Id: id}

	userService := usersvc.New(appEnv.RepoRegistry)
	user, err := userService.FindById(req.Context(), findByIdRequest)

	if err != nil {
		log.WithError(err).Error("Create user failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	if user != nil {
		userResponse := UserResponse(*user)

		_ = appEnv.Render.JSON(w, http.StatusOK, userResponse)
	} else {
		status := common.Error{
			Code:    "NOT_FOUND",
			Message: fmt.Sprintf("User %d is not found", findByIdRequest.Id),
		}

		_ = appEnv.Render.JSON(w, http.StatusNotFound, status)
	}
}

func UpdateUserHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.WithError(err).Error("Bad request")
		return
	}

	updateRequest := usersvc.UpdateRequest{Id: id}

	err = json.Unmarshal(body, &updateRequest)

	if err != nil {
		log.WithError(err).Error("Bad request")
		return
	}

	userService := usersvc.New(appEnv.RepoRegistry)
	user, err := userService.Update(req.Context(), updateRequest)

	if err != nil {
		log.WithError(err).Error("Create user failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	if user != nil {
		userResponse := UserResponse(*user)

		_ = appEnv.Render.JSON(w, http.StatusOK, userResponse)
	} else {
		status := common.Error{
			Code:    "NOT_FOUND",
			Message: fmt.Sprintf("User %d is not found", updateRequest.Id),
		}

		_ = appEnv.Render.JSON(w, http.StatusNotFound, status)
	}
}

func DeleteUserHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	deleteRequest := usersvc.DeleteRequest{Id: id}

	userService := usersvc.New(appEnv.RepoRegistry)

	user, err := userService.Delete(req.Context(), deleteRequest)

	if err != nil {
		log.WithError(err).Error("Delete user failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	if user != nil {
		_ = appEnv.Render.JSON(w, http.StatusNoContent, map[string]string{})
	} else {
		status := common.Error{
			Code:    "NOT_FOUND",
			Message: fmt.Sprintf("User %d is not found", deleteRequest.Id),
		}

		_ = appEnv.Render.JSON(w, http.StatusNotFound, status)
	}
}
