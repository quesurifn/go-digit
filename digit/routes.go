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

func (s *Server) ExchangeToken(c *gin.Context) {
	type ExchangeToken struct {
		PublicToken string `json:"publictoken"`
	}
	var e ExchangeToken
	err := c.BindJSON(&e)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	tok, err := s.Plaid.ExchangeToken(e.PublicToken)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tok)
}
