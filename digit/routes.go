package digit

import (
	"net/http"

	//	gin "gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-gonic/gin"
)

// Serves up Index
func (s *Server) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Go Digit"})
}
