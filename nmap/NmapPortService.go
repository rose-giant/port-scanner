package nmap

import (
	"encoding/xml"
	db "example/goproc/db"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var singlChannelInstance *singlChannel
var singleServiceScanInstance *serviceScanChannel

type State struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

type Service struct {
	Name       string `xml:"name,attr"`
	Mehod      string `xml:"method,attr"`
	Confidence int    `xml:"conf,attr"`
}

type Port struct {
	Protocol string  `xml:"protocol,attr"`
	PortID   int     `xml:"portid,attr"`
	State    State   `xml:"state"`
	Service  Service `xml:"service"`
}

type IPAddress struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Nmaprun struct {
	XMLName xml.Name  `xml:"nmaprun"`
	Ip      IPAddress `xml:"hosthint>address"`
	Ports   []Port    `xml:"host>ports>port"`
}

type singlChannel struct {
	scanChannel chan string
}

type serviceScanChannel struct {
	serviceChannel chan string
}

func GetSingleChannelInstance() *singlChannel {

	if singlChannelInstance == nil {
		singlChannelInstance = &singlChannel{}
		singlChannelInstance.scanChannel = make(chan string, 200048)
	}

	return singlChannelInstance
}

func GetServiceChannelInstance() *serviceScanChannel {

	if singleServiceScanInstance == nil {
		singleServiceScanInstance = &serviceScanChannel{}
		singleServiceScanInstance.serviceChannel = make(chan string, 88888888)
	}

	return singleServiceScanInstance
}

func PortScanServiceFromSingleChannel() {
	fmt.Println("scanning...")
	singleChannel := GetSingleChannelInstance()
	for ch := range singleChannel.scanChannel {
		fmt.Println("sent ip to nmap")
		RunNmapForIp(ch)
	}
}

func RunNmapForIp(ipAddress string) (string, error) {
	fmt.Println("running nmap on the terminal")

	cmd := exec.Command("nmap", "-oX", "-", ipAddress)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error occurred:", err)
		return "", err
	}

	fmt.Println("here's the output: ", out)
	fmt.Println(string(out))
	err = ReadNmapResults(out)
	if err != nil {
		log.Fatal("Error occurred:", err)
		return "", err
	}

	return string(out), nil
}

func runServiceScan(ipAddress string) {
	fmt.Println("running service scan...")
}

func ReadNmapResults(xmlData []byte) error {
	var nmaprun Nmaprun

	fmt.Println("parsing...")
	r := strings.NewReader(string(xmlData))

	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&nmaprun)
	if err != nil {
		log.Println("somthing went wrong while decoding the xml data", err)
		return err
	}

	portsRead2 := PortToPortInUseConvertor(nmaprun)
	fmt.Println("ports read 2: ", portsRead2)
	db.WriteDataToDB(portsRead2, nmaprun.Ip.Addr)

	return nil
}

func WriteIpOnAChannel(ipAddress string) {
	fmt.Println("writing on the channel")
	singleChannel := GetSingleChannelInstance()
	singleChannel.scanChannel <- ipAddress
	fmt.Println("here's what written on the channel: ", ipAddress)
}

func PortToPortInUseConvertor(nmaprun Nmaprun) []db.PortInUse {

	portsRead := []db.PortInUse{}
	var portRead db.PortInUse

	for i := 0; i < len(nmaprun.Ports); i++ {
		portRead.PrtId = nmaprun.Ports[i].PortID
		portRead.Protocol = nmaprun.Ports[i].Protocol
		portRead.State = nmaprun.Ports[i].State.State
		portRead.StateReason = nmaprun.Ports[i].State.Reason
		portRead.ServiceName = nmaprun.Ports[i].Service.Name
		portRead.ServiceMethod = nmaprun.Ports[i].Service.Mehod
		portRead.ServiceConf = nmaprun.Ports[i].Service.Confidence
		portsRead = append(portsRead, portRead)
	}

	return portsRead
}
