package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// usersOptions is a handler for getting the allowed http methods for the users resource.
func usersOptions(c *gin.Context) {
	methods := []string{http.MethodOptions, http.MethodGet, http.MethodPatch, http.MethodHead, http.MethodPost}
	c.Writer.Header().Set("Allow", strings.Join(methods, " "))
	c.String(http.StatusOK, "")
}
