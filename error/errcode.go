package error

var (
	// OK
	ERR_CODE_OK = 0
	ERR_STR_OK  = "ok"

	// SYSTEM ERROR CODE
	ERR_CODE_SYS = 1000
	ERR_STR_SYS  = "system error"
	// TODO: ADD OTHER NEW SYSTEM ERROR, see auth error as example
	ERR_CODE_IO = 1001
	ERR_STR_IO  = "i/o error"

	// DATABASE ERROR CODE
	ERR_CODE_DB = 1100
	ERR_STR_DB  = "db error"
	// TODO: ADD OTHER NEW DATABASE ERROR, see auth error as example

	// NETWORK ERROR CODE
	ERR_CODE_NETWORK = 1200
	ERR_STR_NETWORK  = "network error"
	// TODO: ADD OTHER NEW NETWORK ERROR, see auth error as example

	// PARAMETER ERROR CODE
	ERR_CODE_PARA = 1300
	ERR_STR_PARA  = "parameter error"
	// TODO: ADD OTHER NEW PARAMETER ERROR,

	// AUTH ERROR CODE
	ERR_CODE_AUTH = 1400
	ERR_STR_AUTH  = "auth failed"

	// TODO: ADD OTHER NEW AUTH ERROR, AFTER ERR_CODE_AUTH_EXPIRED AS BELOW
	ERR_CODE_AUTH_EXPIRED = 1401
	ERR_STR_AUTH_EXPIRED  = "auth expired"

	ERR_CODE_CONVERT      = 2000
	ERR_STR_CONVERT       = "convert error"
	ERR_CODE_CONVERT_NIL  = 2001
	ERR_STR_CONVERT_NIL   = "invalid value"
	ERR_CODE_CONVERT_TYPE = 2000
	ERR_STR_CONVERT_TYPE  = "type mismatch"

	// UNKNOWN ERROR
	// ERR_CODE_UNKNOWN = 1900
	ERR_STR_UNKNOWN = "unknown error"
	// TODO: ADD OTHER NEW UNKNOW ERROR, see auth error as example

	errMap = make(map[int]string)
)

func init() {
	// ok
	errMap[ERR_CODE_OK] = ERR_STR_OK

	// system
	errMap[ERR_CODE_SYS] = ERR_STR_SYS
	errMap[ERR_CODE_IO] = ERR_STR_IO

	// database
	errMap[ERR_CODE_DB] = ERR_STR_DB

	// network
	errMap[ERR_CODE_NETWORK] = ERR_STR_NETWORK

	// parameter
	errMap[ERR_CODE_PARA] = ERR_STR_PARA

	// authentication
	errMap[ERR_CODE_AUTH] = ERR_STR_AUTH
	errMap[ERR_CODE_AUTH_EXPIRED] = ERR_STR_AUTH_EXPIRED
	errMap[ERR_CODE_CONVERT] = ERR_STR_CONVERT
	errMap[ERR_CODE_CONVERT_NIL] = ERR_STR_CONVERT_NIL
	errMap[ERR_CODE_CONVERT_TYPE] = ERR_STR_CONVERT_TYPE
}

func fetchErrString(code int) string {
	v, ok := errMap[code]
	if !ok {
		return ERR_STR_UNKNOWN
	}

	return v
}
