package apihandlers

import (
	db "example/goproc/db"
	"example/goproc/nmap"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addreObj struct {
	IP string `json:"ip"`
}

func PostIp(c *gin.Context) {
	fmt.Println("posting your request")
	var newAddress addreObj
	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	}

	nmap.WriteIpOnAChannel(newAddress.IP)
	c.IndentedJSON(http.StatusCreated, newAddress)
}

func GetIp(c *gin.Context) {
	address := c.Param("address")
	res, _ := db.ReadObjectByIpFromdb(address)

	c.JSON(200, gin.H{"ip": res})

}
