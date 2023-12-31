package nmap

import (
	db "example/goproc/db"
	"testing"
)

func Test_getSingleChannelCreatesTheChannel(t *testing.T) {

	channelInstance := GetSingleChannelInstance()
	channelInstance = nil

	channelInstance = GetSingleChannelInstance()
	if channelInstance.scanChannel == nil {
		t.Errorf("the channel is not being made")
	}
}

func Test_getSingleChannelReturnsTheInstance(t *testing.T) {
	expectedInstance := GetSingleChannelInstance()

	if expectedInstance != GetSingleChannelInstance() {
		t.Errorf("The instance is not returned")
	}
}

func Test_getServiceScanReturnsTheInstance(t *testing.T) {
	expectedInstance := GetServiceChannelInstance()

	if expectedInstance != GetServiceChannelInstance() {
		t.Errorf("The instance is not returned")
	}
}

func Test_getServiceScanCreatesTheChannel(t *testing.T) {
	channelInstance := GetServiceChannelInstance()
	channelInstance = nil

	channelInstance = GetServiceChannelInstance()
	if channelInstance.serviceChannel == nil {
		t.Errorf("the channel is not being made")
	}
}

var samplePortId = 1
var sampleProtocol = "tcp"
var sampleServiceName = "s"
var sampleServiceConf = 9
var sampleServiceMethod = "m"
var sampleState = "open"
var sampleStateReason = "reset"
var sampleReasonTTL = ""

func initPort() Port {
	var port Port
	port.PortID = samplePortId
	port.Protocol = sampleProtocol
	port.Service.Name = sampleServiceName
	port.Service.Confidence = sampleServiceConf
	port.Service.Mehod = sampleServiceMethod
	port.State.Reason = sampleStateReason
	port.State.State = sampleState
	port.State.ReasonTTL = sampleReasonTTL

	return port
}

func initPortInUse() db.PortInUse {
	var portInUse db.PortInUse
	portInUse.PrtId = samplePortId
	portInUse.Protocol = sampleProtocol
	portInUse.ServiceConf = sampleServiceConf
	portInUse.ServiceMethod = sampleServiceMethod
	portInUse.ServiceName = sampleServiceName
	portInUse.State = sampleState
	portInUse.StateReason = sampleStateReason

	return portInUse
}

func Test_portToPortInUseConvertorReturnsRightLengthOfPorts(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := initPort()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := initPortInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0] != portsInUse[0] {
		t.Errorf("error while converting")
	}
}

var sampleIP = "127.0.0.1"

func Test_WriteIpOnAChannelWritesSampleIPOnSingleChannel(t *testing.T) {
	WriteIpOnAChannel(sampleIP)
	writtenIP := <-GetSingleChannelInstance().scanChannel

	if writtenIP != sampleIP {
		t.Errorf("error while writing on the channel")
	}
}

// func Test_PortScanServiceFromSingleChannelWalksTheLoop(t *testing.T) {
// 	GetSingleChannelInstance().scanChannel <- sampleIP
// 	PortScanServiceFromSingleChannel()

// }
