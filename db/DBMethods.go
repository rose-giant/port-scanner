package db

import (
	"context"
	"fmt"
	"log"

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

func ReadObjectByIpFromdb(ipAddress string) *PortEmbeddedBesideIP {

	client := GetSingleDBInstance().dbClient
	nmapCollection := client.Database("admin").Collection("nmapResult")
	var result PortEmbeddedBesideIP
	filter := bson.D{{Key: "ip", Value: ipAddress}}

	err := nmapCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("no results found!")
			return nil
		}
		panic(err)
	}

	return &result
}
