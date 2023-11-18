package main

import (
	db "example/goProc/DB"
	"example/goProc/apihandlers"
	"example/goProc/nmap"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_writeIpOnAChannel(t *testing.T) {

	var testCase struct {
		name string
		ip   string
	}

	testCase.name = "write the 82.99.202.35 on the channel"
	testCase.ip = "82.99.202.35"

	t.Run(testCase.name, func(t *testing.T) {
		nmap.WriteIpOnAChannel(testCase.ip)
	})

}

func Test_getSingleChannelInstanceReturnsTheInstance(t *testing.T) {

	var testCase struct {
		name string
	}

	testCase.name = "CheckIfThe Instance Is created"

	t.Run(testCase.name, func(t *testing.T) {
		nmap.GetSingleChannelInstance()
	})
}

func Test_getSingleDBlInstanceReturnsTheInstance(t *testing.T) {
	var testCase struct {
		name string
	}

	testCase.name = "CheckIfThe Instance Is returned"

	t.Run(testCase.name, func(t *testing.T) {
		db.GetSingleDBInstance()
	})
}

func Test_postIpPostsSampleIp(t *testing.T) {
	var testCase struct {
		name string
	}

	t.Run(testCase.name, func(t *testing.T) {
		router := gin.Default()
		router.POST("/addresses", apihandlers.PostIp)
		router.Run("localhost:8080")
	})
}

func Test_runNmapForSampleIP(t *testing.T) {
	var testCase struct {
		name string
		ip   string
	}

	testCase.ip = "82.99.202.35"
	testCase.name = "Check If The IP is run on the nmap service"

	t.Run(testCase.name, func(t *testing.T) {
		nmap.RunNmapForIp(testCase.ip)
	})
}

func Test_readNmapResultsFromFile(t *testing.T) {
	var testCase struct {
		name string
	}

	testCase.name = "Check If The file is read"
	t.Run(testCase.name, func(t *testing.T) {
		db.ReadNmapResultsFromFile()
	})
}
