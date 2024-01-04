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

	log.Init("manual patch new user")
	log.Info("Started manual patch new user")
	datastore.DbInit(config.AppConfig.GetString("databaseUrl"), config.AppConfig.GetString("databaseName"))

	ctx := context.Background()
	err := services.ManualNewAccountInit(ctx, "5f6b9da8cfb9ef7d26ffbcc7")
	if err != nil {
		panic(err)
	}
	log.Info("successfully patch..")

}
