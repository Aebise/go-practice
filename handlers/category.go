package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// categoriesOptions is a handler for getting the allowed http methods for the categories resource.
func categoriesOptions(c *gin.Context) {
	methods := []string{http.MethodOptions, http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPatch}
	c.Writer.Header().Set("Allow", strings.Join(methods, " "))
	c.String(http.StatusOK, "")
}
