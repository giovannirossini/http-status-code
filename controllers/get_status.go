package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giovannirossini/http-status-code/models"
)

// getStatus responds with the list of all HTTP status code as JSON.
func GetStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.HttpStatusCodes)
}

// getStatusById locates the HTTP status code whose ID value matches the id
// parameter sent by the client, then returns that code as a response.
func GetStatusById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of all HTTP status codes, looking for
	// an HTTP status code whose ID value matches the parameter.
	for _, a := range models.HttpStatusCodes {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		} else {
			for _, b := range a.Codes {
				if b.Id == id {
					c.IndentedJSON(http.StatusOK, b)
					return
				}
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "code not found"})
}
