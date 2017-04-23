package digit

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Server is the struct that defines the application server
type Server struct {
	*gin.Engine
}

// NewServer creates a new instance of Server
func NewServer(port string) *Server {
	s := &Server{
		Engine: gin.Default(),
	}
	s.LoadHTMLGlob("templates/*")
	s.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Go Digit",
			},
		)

	})
	return s
}
