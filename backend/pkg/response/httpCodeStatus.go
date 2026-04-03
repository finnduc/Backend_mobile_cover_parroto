package response

const (
	// ===== SUCCESS =====
	CodeSuccess = 200

	// ===== CLIENT ERROR =====
	CodeInvalidParams   = 400
	CodeUnauthorized    = 401
	CodeForbidden       = 403
	CodeNotFound        = 404
	CodeMethodNotAllow  = 405
	CodeRequestTimeout  = 408
	CodeConflict        = 409
	CodeTooManyRequest  = 429

	// ===== BUSINESS ERROR =====
	CodeUserNotFound      = 10001
	CodeUserAlreadyExists = 10002
	CodePasswordWrong     = 10003
	CodeTokenInvalid      = 10004
	CodeTokenExpired      = 10005

	// ===== SERVER ERROR =====
	CodeInternalServerError = 500
	CodeDatabaseError       = 50001
	CodeServiceError        = 50002
	CodeRedisError          = 50003
)

var MsgFlags = map[int]string{

	// success
	CodeSuccess: "Success",

	// client error
	CodeInvalidParams:  "Invalid parameters",
	CodeUnauthorized:   "Unauthorized",
	CodeForbidden:      "Forbidden",
	CodeNotFound:       "Resource not found",
	CodeMethodNotAllow: "Method not allowed",
	CodeRequestTimeout: "Request timeout",
	CodeConflict:       "Conflict",
	CodeTooManyRequest: "Too many requests",

	// business error
	CodeUserNotFound:      "User not found",
	CodeUserAlreadyExists: "User already exists",
	CodePasswordWrong:     "Password incorrect",
	CodeTokenInvalid:      "Token invalid",
	CodeTokenExpired:      "Token expired",

	// server error
	CodeInternalServerError: "Internal server error",
	CodeDatabaseError:       "Database error",
	CodeServiceError:        "Service error",
	CodeRedisError:          "Redis error",
}