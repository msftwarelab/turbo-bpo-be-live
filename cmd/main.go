package main

import (
	"os"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Error("%s", "parameter required: [development,staging,production]")
		os.Exit(1)
	}
	config.Init("graphql", os.Args[1])

	datastore.DbInit(config.AppConfig.GetString("databaseUrl"), config.AppConfig.GetString("databaseName"))
	awsS3.Init()
	var srv = Server{}
	srv.start()
}
