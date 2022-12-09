package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}
func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func Forget(c *gin.Context, message, tip string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip":     tip,
	})
}
func Comment(c *gin.Context, floor int, comment string) {
	c.JSON(http.StatusOK, gin.H{
		"floor":   floor,
		"comment": comment,
	})
}
