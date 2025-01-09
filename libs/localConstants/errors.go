package localconstants

const (
	SERVER_ERROR        = "server error"
	UNAUTHORIZED        = "unauthorized"
	FORBIDDEN           = "forbidden"
	INVALID_JSON_FORMAT = "invalid Json format"
)

var ErrorMap = map[int]string{
	500: SERVER_ERROR,
	403: FORBIDDEN,
	401: UNAUTHORIZED,
}
