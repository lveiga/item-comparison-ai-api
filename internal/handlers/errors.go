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
	ErrInvalidProductID       = NewError(http.StatusBadRequest, "Invalid product ID")
	ErrFailedToLoadProducts   = NewError(http.StatusInternalServerError, "Failed to load products")
	ErrProductNotFound        = NewError(http.StatusNotFound, "Product not found")
	ErrInvalidLimitParameter  = NewError(http.StatusBadRequest, "Invalid limit parameter")
	ErrInvalidOffsetParameter = NewError(http.StatusBadRequest, "Invalid offset parameter")
	ErrFailedToSaveProducts   = NewError(http.StatusInternalServerError, "Failed to save products")
	ErrBindJSON               = NewError(http.StatusBadRequest, "Invalid request body")
)

// HandleError sends an error response.
func HandleError(c *gin.Context, err *Error) {
	c.JSON(err.Code, gin.H{"error": err.Message})
}
