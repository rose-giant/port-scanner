package nmap

import (
	db "example/goProc/DB"
	"fmt"
	"log"
	"os/exec"
)

var singlChannelInstance *singlChannel

type addreObj struct {
	IP string `json:"ip"`
}

type singlChannel struct {
	scanChannel chan string
}

func GetSingleChannelInstance() *singlChannel {

	if singlChannelInstance == nil {
		singlChannelInstance = &singlChannel{}
		singlChannelInstance.scanChannel = make(chan string, 200048)
	}

	return singlChannelInstance
}

func PortScanServiceFromSingleChannel() {
	fmt.Println("scanning...")
	singleChannel := GetSingleChannelInstance()
	for ch := range singleChannel.scanChannel {
		RunNmapForIp(ch)
	}
}

func RunNmapForIp(ipAddress string) {
	fmt.Println("running nmap on the terminal")
	cmd := exec.Command("nmap", "-oX", "./scanResult.xml", ipAddress)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("error occured:", err)
	}

	fmt.Printf("%s", out)
	db.ReadNmapResultsFromFile()
}

func WriteIpOnAChannel(ipAddress string) {
	fmt.Println("writing on the channel")
	singleChannel := GetSingleChannelInstance()
	singleChannel.scanChannel <- ipAddress
	fmt.Println("here's what written on the channel: ", ipAddress)
}
