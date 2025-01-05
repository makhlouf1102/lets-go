package commonerrors

import (
	localconstants "lets-go/libs/localConstants"
	loglib "lets-go/libs/logLib"
	"net/http"
)

func EncodingError(responseWriter http.ResponseWriter, err error) {
	loglib.LogError("error while encoding response", err)
	http.Error(responseWriter, localconstants.SERVER_ERROR, http.StatusInternalServerError)
}

func DencodingError(responseWriter http.ResponseWriter, err error) {
	http.Error(responseWriter, localconstants.INVALID_JSON_FORMAT, http.StatusBadRequest)
}
