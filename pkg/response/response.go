package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessReponseData sends a 200 HTTP response with internal code + data
func SuccessReponseData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    CodeSuccess,
		Message: MsgFlags[CodeSuccess],
		Data:    data,
	})
}

// ErrorResponseData sends a 200 HTTP response with error code (frontend reads code field)
func ErrorResponseData(c *gin.Context, code int, data interface{}) {
	msg, ok := MsgFlags[code]
	if !ok {
		msg = "unknown error"
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
