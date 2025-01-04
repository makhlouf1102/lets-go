package responsewriter

import (
	loglib "lets-go/libs/logLib"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, err error, logText string, httpText string, httpStatus int) {
	if err != nil {
		loglib.LogError(logText, err)
		http.Error(w, httpText, httpStatus)
	}
}
