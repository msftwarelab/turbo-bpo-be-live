package main

import (
	"context"
	"os"

	"github.com/lonmarsDev/bpo-golang-grahpql/internal/services"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func main() {

	if len(os.Args) != 2 {
		log.Error("%s", "parameter required: [development,staging,production]")
		os.Exit(1)
	}
	config.Init("jobs", os.Args[1])

	log.Init("batch update pipeline notes")
	log.Info("Started update pipeline notes")
	datastore.DbInit(config.AppConfig.GetString("databaseUrl"), config.AppConfig.GetString("databaseName"))

	ctx := context.Background()
	services.BatchUpdatePipelineNotes(ctx)
	log.Info("done update pipeline notes")

}
