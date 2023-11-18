package db

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var singleDBInstance *singleDB

type singleDB struct {
	dbClient *mongo.Client
}

type PortInUse struct {
	Protocol      string
	PrtId         int
	State         string
	StateReason   string
	ServiceName   string
	ServiceMethod string
	ServiceConf   int
}

type PortEmbeddedBesideIP struct {
	IP    string
	Ports []PortInUse
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

type Nmaprun struct {
	XMLName xml.Name  `xml:"nmaprun"`
	Ip      IPAddress `xml:"hosthint>address"`
	Ports   []Port    `xml:"host>ports>port"`
}

func GetSingleDBInstance() *singleDB {
	if singleDBInstance == nil {
		singleDBInstance = &singleDB{}
	}

	return singleDBInstance
}

func StartDBConnection() {
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	fmt.Println("db connection started")
	GetSingleDBInstance().dbClient = client
}

func WriteDataToDB(ports []PortInUse, ipAddress string) {
	client := GetSingleDBInstance().dbClient
	nmapCollection := client.Database("admin").Collection("nmapResult")
	var toBeWrittenTOdb PortEmbeddedBesideIP
	toBeWrittenTOdb.Ports = ports
	toBeWrittenTOdb.IP = ipAddress
	result, err := nmapCollection.InsertOne(context.TODO(), toBeWrittenTOdb)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}

func ReadNmapResultsFromFile() {
	var nmaprun Nmaprun
	file, err := os.Open("scanResult.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&nmaprun)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	//fmt.Println("\n", nmaprun.Ports, "\n", nmaprun.Ip.Addr, "\n")
	portsRead := []PortInUse{}
	var portRead PortInUse
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

	WriteDataToDB(portsRead, nmaprun.Ip.Addr)
}

func ReadObjectByIpFromdb(ipAddress string) string {

	client := GetSingleDBInstance().dbClient
	nmapCollection := client.Database("admin").Collection("nmapResult")
	var result PortEmbeddedBesideIP
	filter := bson.D{{Key: "ip", Value: ipAddress}}

	err := nmapCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	if len(result.Ports) < 1 {
		log.Println("no results found!")
		return ""
	}

	res, _ := json.Marshal(result)
	return string(res)
}