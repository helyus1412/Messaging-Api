package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {
	mapUrls()

	port := os.Getenv("APP_PORT")
	router.Run(port)
}
