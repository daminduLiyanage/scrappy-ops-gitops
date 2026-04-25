package main

import (
	"log"
	"net/http"
	"scrappy-backend/internal/scraper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// In RIBs, this initialization happens in a 'Builder'
	interactor := scraper.NewInteractor()

	r.GET("/hello", func(c *gin.Context) {
		res, err := interactor.ExecuteHello("Sri Lanka Real Estate Market")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res})
	})

	log.Println("Scrappy Backend running on :8080")
	r.Run(":8080")
}
