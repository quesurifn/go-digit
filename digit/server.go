package digit

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/quesurifn/go-digit/banker"
)

// Server is the struct that defines the application server
type Server struct {
	*gin.Engine
	port  string
	Plaid *banker.Plaid
}

// NewServer creates a new instance of Server
func NewServer(port string) *Server {
	return &Server{
		Engine: gin.Default(),
		port:   port,
	}
}

// Runs gin
func (s *Server) Run() error {
	s.configure()
	return s.Engine.Run(fmt.Sprintf(":%s", s.port))
}

// Sets logger and recovery
func (s *Server) configure() {
	s.Use(gin.Logger())
	s.Use(gin.Recovery())
	s.Use(static.Serve("/", static.LocalFile("./templates/open-digit/dist/", true)))
	//	s.Static("/", "./templates/open-digit/dist/")
	s.LoadHTMLGlob("templates/open-digit/dist/*.html")
	s.GET("/", s.Index)
	s.POST("/plaid/exchangeToken", s.ExchangeToken)
}
