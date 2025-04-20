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
	loglib.LogError("error while decoding response", err)
	http.Error(responseWriter, localconstants.INVALID_JSON_FORMAT, http.StatusBadRequest)
}

func HttpError(responseWriter http.ResponseWriter, code int) {
	http.Error(responseWriter, localconstants.ErrorMap[code], code)
}

func HttpErrorWithMessage(responseWriter http.ResponseWriter, err error, code int, message string) {
	loglib.LogError(message, err)
	http.Error(responseWriter, localconstants.ErrorMap[code], code)
}