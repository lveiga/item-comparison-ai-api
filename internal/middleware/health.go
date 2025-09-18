package middlewares

import (
	"net/http"

	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"

	"github.com/gin-gonic/gin"
)

// Health ....
func Health(db *database.Database, c *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := db.CheckLiveness(c.DatabasePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":     http.StatusInternalServerError,
				"database": "DOWN",
				"app":      "DOWN",
				"error":    err,
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"database": "UP",
			"app":      "UP",
		})
		return
	}
}
