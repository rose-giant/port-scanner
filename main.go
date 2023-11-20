package main

import (
	db "example/goProc/DB"
	"example/goProc/apihandlers"
	"example/goProc/nmap"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("started")

	go nmap.PortScanServiceFromSingleChannel()
	go db.StartDBConnection()

	router := gin.Default()
	router.GET("/addresses/:address", apihandlers.GetIp)
	router.POST("/addresses", apihandlers.PostIp)
	router.Run("localhost:8080")
}
