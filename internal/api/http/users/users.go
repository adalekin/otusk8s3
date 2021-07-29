package users

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/appenv"
)

func ListUsersHandler(w http.ResponseWriter, req *http.Request, appEnv appenv.AppEnv) {
	userRepository := appEnv.RepoRegistry.GetUserRepository()
	users, err := userRepository.FindAll(req.Context())

	if err != nil {
		log.WithError(err).Error("Transaction failed")

		status := common.Error{
			Code:    "ERROR",
			Message: fmt.Sprint(err),
		}

		_ = appEnv.Render.JSON(w, http.StatusInternalServerError, status)
		return
	}

	_ = appEnv.Render.JSON(w, http.StatusOK, users)
}
