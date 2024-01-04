package main

import (
	"context"
	"github.com/lonmarsDev/bpo-golang-grahpql/internal/services"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Error("%s", "parameter required: [development,staging,production]")
		os.Exit(1)
	}
	config.Init("jobs", os.Args[1])

	log.Init("batch update invoice")
	log.Info("Started update invoice")
	datastore.DbInit(config.AppConfig.GetString("databaseUrl"), config.AppConfig.GetString("databaseName"))

	ctx := context.Background()
	services.BatchUpdateInvoice(ctx)
	log.Info("done update invoice")

}
