package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponseFormat struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type ErrorResponseFormat struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
    response := SuccessResponseFormat{
        Status  :  "success",
        Message : message,
        Data    : data,
    }

    c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
    response := ErrorResponseFormat{
        Status  : "error",
        Message : message,
        Data    : nil,
    }

    c.JSON(statusCode, response)
}
