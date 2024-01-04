package datastore

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DbCollections = Collections{}
	mutex         = &sync.Mutex{}
)

type Collections struct {
	Users                          *mongo.Collection
	ProfileDoc                     *mongo.Collection
	Accounts                       *mongo.Collection
	Adjustments                    *mongo.Collection
	Comments                       *mongo.Collection
	Defaults                       *mongo.Collection
	Pipelines                      *mongo.Collection
	PipeLineQualityControlAndNotes *mongo.Collection
	PipeLineDocs                   *mongo.Collection
	PipeLinePhotos                 *mongo.Collection
	PipeLineNotes                  *mongo.Collection
	Credits                        *mongo.Collection
	Companies                      *mongo.Collection
	PipelineStates                 *mongo.Collection
	EmailTemplates                 *mongo.Collection
	Headers                        *mongo.Collection
	Instructions                   *mongo.Collection
	QualityControls                *mongo.Collection
	PipeLineNeighborhood           *mongo.Collection
	Reviews                        *mongo.Collection
	Requests                       *mongo.Collection
	RequestsHistory                *mongo.Collection
	Invoices                       *mongo.Collection
	Iforms                         *mongo.Collection
	IformTemps                     *mongo.Collection
	PipeLineRepairs                *mongo.Collection
	Announcements                  *mongo.Collection
	Sessions                       *mongo.Collection
	PermissionGroups               *mongo.Collection
	LoginLogs                      *mongo.Collection
	PipelineComparables            *mongo.Collection
	Billings                       *mongo.Collection
	IformGrids                     *mongo.Collection
	CreditLedgers                  *mongo.Collection
}

func DbInit(dbUrl, dbName string) {
	mutex.Lock()
	defer mutex.Unlock()
	database := CreateDbConnection(dbUrl, dbName)
	DbCollections = Collections{
		Users:                          database.Collection("users"),
		ProfileDoc:                     database.Collection("profileDoc"),
		Accounts:                       database.Collection("accounts"),
		Adjustments:                    database.Collection("adjustments"),
		Comments:                       database.Collection("comments"),
		Defaults:                       database.Collection("defaults"),
		Pipelines:                      database.Collection("pipeLines"),
		PipeLineQualityControlAndNotes: database.Collection("pipeLineQualityControlAndNotes"),
		PipeLineDocs:                   database.Collection("pipeLineDocs"),
		PipeLinePhotos:                 database.Collection("pipeLinePhotos"),
		PipeLineNotes:                  database.Collection("pipeLineQualityControlAndNotes"),
		Credits:                        database.Collection("credits"),
		Companies:                      database.Collection("companies"),
		PipelineStates:                 database.Collection("pipelineStates"),
		EmailTemplates:                 database.Collection("emailTemplates"),
		Headers:                        database.Collection("headers"),
		Instructions:                   database.Collection("instructions"),
		QualityControls:                database.Collection("qualityControls"),
		PipeLineNeighborhood:           database.Collection("pipeLineNeighborhood"),
		Reviews:                        database.Collection("reviews"),
		Requests:                       database.Collection("requests"),
		RequestsHistory:                database.Collection("requestsHistory"),
		Invoices:                       database.Collection("invoices"),
		Iforms:                         database.Collection("iforms"),
		IformTemps:                     database.Collection("iformTemps"),
		PipeLineRepairs:                database.Collection("pipelineRepairs"),
		Announcements:                  database.Collection("announcements"),
		Sessions:                       database.Collection("sessions"),
		PermissionGroups:               database.Collection("permissionGroups"),
		LoginLogs:                      database.Collection("loginLogs"),
		PipelineComparables:            database.Collection("pipelineComparables"),
		Billings:                       database.Collection("billings"),
		IformGrids:                     database.Collection("iformGrids"),
		CreditLedgers:                  database.Collection("creditLedgers"),
	}
}

func CreateDbConnection(url string, database string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client.Database(database)
}
