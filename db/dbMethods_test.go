package db

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = "admin"
var CollectionName = "nmapResult"

func Test_getSingleDbInstanceReturnsInstance(t *testing.T) {
	dbInstance := GetSingleDBInstance()

	if dbInstance != GetSingleDBInstance() {
		t.Errorf("Error while returning the db instance")
	}
}

func Test_getSingleDbInstanceCreatesNewInstance(t *testing.T) {
	dbInstance := GetSingleDBInstance()
	dbInstance = nil
	dbInstance = GetSingleDBInstance()
	if dbInstance == nil {
		t.Errorf("Error while creating a new DB instance")
	}
}

func Test_startDbConnectionCreatesTheConnectionWithCorrectLog(t *testing.T) {
	err := StartDBConnection()
	if err != nil {
		t.Errorf("Error while creating the connection")
	}
}

// func Test_startsDbConnectionCreatesValidClient(t *testing.T) {
// 	StartDBConnection()
// 	client := GetSingleDBInstance().dbClient
// 	err := client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Errorf("Error while validating the client")
// 	}
// }

var samplePortId = 1
var sampleProtocol = "tcp"
var sampleServiceName = "s"
var sampleServiceConf = 9
var sampleServiceMethod = "m"
var sampleState = "open"
var sampleStateReason = "reset"

var sampleIPAddress = "127.0.0.1"

func initPortInUse() PortInUse {
	var port PortInUse
	port.PrtId = samplePortId
	port.Protocol = sampleProtocol
	port.ServiceName = sampleServiceName
	port.ServiceConf = sampleServiceConf
	port.ServiceMethod = sampleServiceMethod
	port.StateReason = sampleStateReason
	port.State = sampleState

	return port
}

func disconnectDBMock() {
	err := mtest.Teardown()
	if err != nil {
		fmt.Printf("tear down error: %s", err)
	}
}

func initMockDBConnection() error {
	err2 := mtest.Setup(mtest.NewSetupOptions())
	if err2 != nil {
		return fmt.Errorf("error while setting up %s", err2)
	}

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(mtest.ClusterURI()))

	if err != nil {
		return fmt.Errorf("error while connecting %s", err)
	}

	fmt.Println("db connection started")
	GetSingleDBInstance().dbClient = client
	return nil
}

func Test_startDBConnectionStartsTheConnection(t *testing.T) {
	err := StartDBConnection()
	if err != nil {
		fmt.Println("db connection error: ", err)
		t.Errorf("Connection failed")
	}
}

func Test_writeToDBWritesPortsInUseToMockDB(t *testing.T) {
	err := initMockDBConnection()
	if err != nil {
		fmt.Printf("error while connecting %s", err)
		return
	}

	defer disconnectDBMock()

	portInUse := initPortInUse()
	var portsInUse = []PortInUse{}
	portsInUse = append(portsInUse, portInUse)

	WriteDataToDB(portsInUse, sampleIPAddress)

	var result PortEmbeddedBesideIP
	filter := bson.D{{Key: "ip", Value: sampleIPAddress}, {Key: "ports", Value: portsInUse}}
	client := GetSingleDBInstance().dbClient
	dbCollection := client.Database(DBName).Collection(CollectionName)
	err = dbCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Error while writing to db")
	}
}

var findByIpVal = "82.99.202.35"

func Test_readFromDbReadsFromMockDb(t *testing.T) {
	err := initMockDBConnection()
	if err != nil {
		return
	}

	defer disconnectDBMock()

	result, err := ReadObjectByIpFromdb(findByIpVal)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Error while reading")
	}

	fmt.Println(result)
}
