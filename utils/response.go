package utils

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func SuccessResponse(c *gin.Context, data interface{}, message string) {
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{
        "status":  "error",
        "message": message,
        "data":    nil,
    })
}
