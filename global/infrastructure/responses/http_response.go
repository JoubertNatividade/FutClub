package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func Conflict(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{
		"error": http.StatusText(http.StatusConflict),
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": http.StatusText(http.StatusNotFound),
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": http.StatusText(http.StatusInternalServerError),
	})
}

func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": http.StatusText(http.StatusCreated),
	})
}

func Success(c *gin.Context, data any) {
	if data == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, http.StatusText(http.StatusNoContent))
}
