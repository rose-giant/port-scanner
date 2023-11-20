package nmap

import (
	db "example/goProc/DB"
	"testing"
)

func TestDivide(t *testing.T) {

	expected := 2.0
	got := Divide(10.5, 5.25)

	if got != expected {
		t.Errorf("expected %.1f , got %.1f", expected, got)
	}
}

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

var samplePortId = 1
var sampleProtocol = "tcp"
var sampleServiceName = "s"
var sampleServiceConf = 9
var sampleServiceMethod = "m"
var sampleState = "open"
var sampleStateReason = "reset"
var sampleReasonTTL = ""

func init_port() Port {
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

func init_portInUse() db.PortInUse {
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
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if len(PortToPortInUseConvertor(nmaprun)) != len(portsInUse) {
		t.Errorf("error while converting")
	}

}

func Test_portToPortInUseConvertorAssignsRightProtocol(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].Protocol != portsInUse[0].Protocol {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightPortId(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].PrtId != portsInUse[0].PrtId {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightState(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].State != portsInUse[0].State {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightStateReason(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].StateReason != portsInUse[0].StateReason {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightStateServiceName(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].ServiceName != portsInUse[0].ServiceName {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightStateServiceMethod(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].ServiceMethod != portsInUse[0].ServiceMethod {
		t.Errorf("error while converting")
	}
}

func Test_portToPortInUseConvertorAssignsRightStateServiceConf(t *testing.T) {

	var nmaprun Nmaprun
	var ports []Port
	port := init_port()
	ports = append(ports, port)
	nmaprun.Ports = ports

	var portsInUse []db.PortInUse
	portInUse := init_portInUse()
	portsInUse = append(portsInUse, portInUse)

	if PortToPortInUseConvertor(nmaprun)[0].ServiceConf != portsInUse[0].ServiceConf {
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