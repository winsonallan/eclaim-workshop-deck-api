package middleware

import (
	"fmt"
	"net/http"

	"eclaim-workshop-deck-api/internal/common/response"

	"github.com/gin-gonic/gin"
)

type UploadConfig struct {
	MaxSizeBytes  int64
	AllowedTypes  map[string]bool
	FormFieldName string // the multipart field name, e.g. "file"
}

func ValidateUpload(cfg UploadConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile(cfg.FormFieldName)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "no file provided")
			c.Abort()
			return
		}
		defer file.Close()

		// Check size
		if header.Size > cfg.MaxSizeBytes {
			response.Error(c, http.StatusBadRequest,
				fmt.Sprintf("file too large, max %dMB", cfg.MaxSizeBytes/1024/1024),
			)
			c.Abort()
			return
		}

		// Check content type
		contentType := header.Header.Get("Content-Type")
		if !cfg.AllowedTypes[contentType] {
			response.Error(c, http.StatusBadRequest, "file type not allowed")
			c.Abort()
			return
		}

		c.Next()
	}
}
