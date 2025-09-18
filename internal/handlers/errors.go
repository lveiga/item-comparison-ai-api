package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error represents a handler error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError creates a new Error instance.
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Error messages
var (
	ErrInvalidID       = NewError(http.StatusBadRequest, "Invalid ID")
	ErrFailedToLoad   = NewError(http.StatusInternalServerError, "Failed to load")
	ErrNotFound        = NewError(http.StatusNotFound, "Not found")
	ErrInvalidLimitParameter  = NewError(http.StatusBadRequest, "Invalid limit parameter")
	ErrInvalidOffsetParameter = NewError(http.StatusBadRequest, "Invalid offset parameter")
	ErrFailedToSave   = NewError(http.StatusInternalServerError, "Failed to save")
	ErrBindJSON               = NewError(http.StatusBadRequest, "Invalid request body")
)

// HandleError sends an error response.
func HandleError(c *gin.Context, err *Error) {
	c.JSON(err.Code, gin.H{"error": err.Message})
}
