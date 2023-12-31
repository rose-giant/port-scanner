package db

import (
	"context"
	"fmt"

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

func GetSingleDBInstance() *singleDB {
	if singleDBInstance == nil {
		singleDBInstance = &singleDB{}
	}

	return singleDBInstance
}

func StartDBConnection() error {
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	fmt.Println("db connection started")
	GetSingleDBInstance().dbClient = client
	return nil
}

func WriteDataToDB(ports []PortInUse, ipAddress string) error {
	client := GetSingleDBInstance().dbClient
	nmapCollection := client.Database("admin").Collection("nmapResult")
	var toBeWrittenTOdb PortEmbeddedBesideIP
	toBeWrittenTOdb.Ports = ports
	toBeWrittenTOdb.IP = ipAddress
	result, err := nmapCollection.InsertOne(context.TODO(), toBeWrittenTOdb)
	fmt.Println(result)

	if err != nil {
		return err
	}

	return nil
}

func ReadObjectByIpFromdb(ipAddress string) (*PortEmbeddedBesideIP, error) {

	client := GetSingleDBInstance().dbClient
	nmapCollection := client.Database("admin").Collection("nmapResult")
	var result PortEmbeddedBesideIP
	filter := bson.D{{Key: "ip", Value: ipAddress}}

	err := nmapCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("error while finding: ", err)
		return nil, err
	}

	return &result, nil
}
