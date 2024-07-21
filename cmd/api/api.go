package main

// Gin boilerplate with ping endpoint

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// PING - for readiness and liveness probes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Version
	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v0.0.1",
		})
	})

	// Get the status of the cluster
	r.GET("/cluster/:name", func(c *gin.Context) {
		result, err := ping(c.Param("name"))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": result,
		})
	})

	// Create a new cluster
	r.POST("/cluster", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)
	})

	// Delete a cluster
	r.DELETE("/cluster/:name", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)
	})

	// Run the server
	r.Run()
}

// ping will send a ping to the name of the cluster and return the response
func ping(name string) (string, error) {

	// Launch a ping command
	log.Printf("Pinging %s\n", name)
	ec := exec.Command("ping", "-c", "3", name)
	if err := ec.Run(); err != nil {
		return "fail", fmt.Errorf("ERROR running escrowAnalyzer: %s", err)
	}

	// and return the response
	return "success", nil
}
