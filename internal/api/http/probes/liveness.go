package probes

import (
	"net/http"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/appenv"
)

func LivenessHandler(appEnv appenv.AppEnv, w http.ResponseWriter, _ *http.Request) {
	status := common.Error{
		Code: "OK",
	}

	appEnv.Render.JSON(w, http.StatusOK, status)
}
