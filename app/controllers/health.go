package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetHelthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
