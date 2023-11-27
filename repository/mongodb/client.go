package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	osx "github.com/garagestandard/roomap-lib-golang/os"
)

const (
	connectionStringTemplate = "mongodb://%s:%s@%s:%s/%s?tls=true&replicaSet=rs0&readpreference=%s"
)

var err error
var client *mongo.Client

func InitClient() {

	caFilePath := osx.Getenv("MONGODB_CA_FILE_PATH", "global-bundle.pem")
	dbu := osx.Getenv("MONGODB_USERNAME", "")
	dbpw := osx.Getenv("MONGODB_PASSWORD", "")
	dbh := osx.Getenv("MONGODB_ENDPOINT", "")
	dbpt := osx.Getenv("MONGODB_PORT", "27017")
	dbn := osx.Getenv("MONGODB_NAME", "")
	dbrp := osx.Getenv("MONGODB_READ_PREFERENCE", "secondaryPreferred")

	dbct := osx.Getenv("MONGODB_CONNECTION_TIMEOUT", "5")
	dbqt := osx.Getenv("MONGODB_QUERY_TIMEOUT", "30")

	dbconf := fmt.Sprintf(connectionStringTemplate,
		dbu, dbpw, dbh, dbpt, dbn, dbrp)

	tlsConfig, err := getCustomTLSConfig(caFilePath)
	if err != nil {
		log.Fatalf("Failed getting TLS configuration: %v", err)
	}

	connectionTimeout, _ := strconv.Atoi(dbct)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectionTimeout)*time.Second)
	defer cancel()

	queryTimeout, _ := strconv.Atoi(dbqt)
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(queryTimeout)*time.Second)
	defer cancel()

	client, err = mongo.NewClient(options.Client().ApplyURI(dbconf).SetTLSConfig(tlsConfig))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	log.Println("Connected to MongoDB!")
	log.Printf("- Database: <%s>\n", dbn)
	log.Printf("- Connection Timeout:<%s>\n", dbct)
	log.Printf("- Query Timeout:<%s>\n", dbqt)
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}
	return tlsConfig, nil
}

func GetClient() *mongo.Client {

	if client == nil {
		InitClient()
	}
	return client
}

func GetDatabase() *mongo.Database {
	dbn := osx.Getenv("MONGODB_NAME", "")
	return GetClient().Database(dbn)
}
