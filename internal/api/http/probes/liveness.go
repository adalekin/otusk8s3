package probes

import (
	"net/http"

	"github.com/adalekin/otusk8s3/internal/api/http/common"
	"github.com/adalekin/otusk8s3/internal/appenv"
)

func LivenessHandler(w http.ResponseWriter, _ *http.Request, appEnv appenv.AppEnv) {
	status := common.Error{
		Code: "OK",
	}

	appEnv.Render.JSON(w, http.StatusOK, status)
}
