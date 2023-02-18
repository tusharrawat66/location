package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{"ID": id})
}
