package main

import (
	"context"
	"os"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

type NewPipeline struct {
	NewStatus      string
	NewOrderAmount *float64
}

func main() {

	if len(os.Args) != 2 {
		log.Error("%s", "parameter required: [development,staging,production]")
		os.Exit(1)
	}
	config.Init("jobs", os.Args[1])

	log.Init("order-timer-job")
	log.Info("Started order timer job")
	datastore.DbInit(config.AppConfig.GetString("databaseUrl"), config.AppConfig.GetString("databaseName"))
	//run jobs every 60 sec

	frequency := time.Second * time.Duration(60)
	ctx, _ := context.WithCancel(context.Background())

	for {

		var pipelineStatuses = make(map[string]NewPipeline)

		filterData := datastore.SearchFilterByStatuses([]string{constants.PipelineStatusActive, constants.PipelineStatusStandBy, constants.PipelineStatusLate})
		pipelinesRaw, _ := datastore.SearchPipelines(ctx, filterData, 0, 100)
		pipelineStatuses = timerCondition(ctx, pipelinesRaw)

		updateNewStatuses(ctx, pipelineStatuses)
		time.Sleep(frequency)
	}
}

func updateNewStatuses(ctx context.Context, pipelineStatuses map[string]NewPipeline) {

	var pipelineIDsForCompleted []string
	var pipelineIDsForCancelled []string
	var pipelineIDsForLate []string
	for key, value := range pipelineStatuses {
		if value.NewStatus == constants.PipelineStatusComplete {
			pipelineIDsForCompleted = append(pipelineIDsForCompleted, key)
		}
		if value.NewStatus == constants.PipelineStatusCancelled {
			pipelineIDsForCancelled = append(pipelineIDsForCancelled, key)
		}
		if value.NewStatus == constants.PipelineStatusLate {
			pipelineIDsForLate = append(pipelineIDsForLate, key)
		}
	}
	if len(pipelineIDsForCompleted) > 0 {
		filterForComplete := datastore.FilterByIds(pipelineIDsForCompleted)
		isUpdateSuccess, err := datastore.UpdatePipelineStatuses(ctx, filterForComplete, constants.PipelineStatusComplete, nil)
		if err != nil {
			log.Error("error on update status to complete , error : %d on ids: %s", err, datastore.PrettyPrint(pipelineIDsForCompleted))
		}
		if isUpdateSuccess {
			log.Info("Update success on status complete ids :%s", datastore.PrettyPrint(pipelineIDsForCompleted))
		}
	}

	if len(pipelineIDsForCancelled) > 0 {
		filterForCancelled := datastore.FilterByIds(pipelineIDsForCancelled)
		isUpdateSuccess, err := datastore.UpdatePipelineStatuses(ctx, filterForCancelled, constants.PipelineStatusCancelled, nil)
		if err != nil {
			log.Error("error on update status to cancelled , error : %d on ids: %s", err, datastore.PrettyPrint(pipelineIDsForCompleted))
		}
		if isUpdateSuccess {
			log.Info("Update success on status complete ids :%s", datastore.PrettyPrint(pipelineIDsForCompleted))
		}
	}

	if len(pipelineIDsForLate) > 0 {

		filterForLate := datastore.FilterByIds(pipelineIDsForLate)
		isUpdateSuccess, err := datastore.UpdatePipelineStatuses(ctx, filterForLate, constants.PipelineStatusLate, nil)
		if err != nil {
			log.Error("error on update status to complete , error : %d on ids: %s", err, datastore.PrettyPrint(pipelineIDsForLate))
		}
		if isUpdateSuccess {
			log.Info("Update success on status complete ids :%s", datastore.PrettyPrint(pipelineIDsForLate))
		}

		// update orderfee
		for k, v := range pipelineStatuses {
			if v.NewOrderAmount != nil {
				if *v.NewOrderAmount > 0 {
					filterForLate := datastore.FilterByIds([]string{k})
					isUpdateSuccess, err := datastore.UpdatePipelineStatuses(ctx, filterForLate, constants.PipelineStatusLate, v.NewOrderAmount)
					if err != nil {
						log.Error("error on update status to complete , error : %d on ids: %s", err, datastore.PrettyPrint(pipelineIDsForLate))
					}
					if isUpdateSuccess {
						log.Info("Update orderfee complete id :%s new fee amount : %s", datastore.PrettyPrint(pipelineIDsForLate), datastore.PrettyPrint(v.NewOrderAmount))
					}

				}
			}
		}
	}

}

func timerCondition(ctx context.Context, pipelineRaws []*datastore.Pipeline) map[string]NewPipeline {

	pipelineState, _ := datastore.GetPipelineState(ctx)

	var pipelineStatuses = make(map[string]NewPipeline)
	for _, v := range pipelineRaws {

		if v.DueDateTime != nil {
			duedatetime := pointers.PrimativeToDateTime(*v.DueDateTime)

			if strings.ObjectTOString(v.Status) == constants.PipelineStatusActive && duedatetime.Before(time.Now().Local()) {

				pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, nil}

				if pointers.ObjectTOBool(v.IsSuperRush) {
					newOrderAmount := *v.OrderFee - float64(*pipelineState.OPSuperRush)
					pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, pointers.Float64(newOrderAmount)}
				}

				if pointers.ObjectTOBool(v.IsRushOrder) {
					newOrderAmount := *v.OrderFee - float64(*pipelineState.OPRush)
					pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, pointers.Float64(newOrderAmount)}
				}

				// // condition on 2days completed status
				// // if status is standby
				// // if order equal to interior
				// //set to complete
				// activationDatePlus48hrs := pointers.PrimativeToDateTime(*v.ActivationDateTime).Add(time.Hour * time.Duration(48))
				// if goString.ToUpper(strings.ObjectTOString(v.OrderType)) == constants.PipelineOrderTypeInterior && strings.ObjectTOString(v.Status) == constants.PipelineStatusStandBy && activationDatePlus48hrs.Before(time.Now().Local()) {
				// 	pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusComplete, nil}
				// }

				// // //if status is active , if 2days from activation
				// // // if order type not equal interior
				// // // set to cancelled
				// // if goString.ToUpper(strings.ObjectTOString(v.OrderType)) != constants.PipelineOrderTypeInterior && strings.ObjectTOString(v.Status) == constants.PipelineStatusActive && activationDatePlus48hrs.Before(time.Now().Local()) {
				// // 	pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusCancelled, nil}
				// // }

				// //if 24hrs from active set to late
				// // if order type not equal to interior
				// //set to late
				// activationDateTimePlus24 := pointers.PrimativeToDateTime(*v.ActivationDateTime).Add(time.Hour * time.Duration(24))
				// if goString.ToUpper(strings.ObjectTOString(v.OrderType)) != constants.PipelineOrderTypeInterior && strings.ObjectTOString(v.Status) == constants.PipelineStatusActive && activationDateTimePlus24.Before(time.Now().Local()) {
				// 	pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, nil}
				// }

				// //condition if Super Rush (2 Hrs)
				// //if and status is active
				// //set to late
				// if v.IsSuperRush != nil {
				// 	activationDateTimePlus2hrs := pointers.PrimativeToDateTime(*v.ActivationDateTime).Add(time.Hour * time.Duration(6))
				// 	if strings.ObjectTOString(v.Status) == constants.PipelineStatusActive && activationDateTimePlus2hrs.Before(time.Now().Local()) && *v.IsSuperRush {
				// 		newOrderAmount := *v.OrderFee - float64(*pipelineState.OPSuperRush)
				// 		pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, pointers.Float64(newOrderAmount)}
				// 	}
				// }

				// //condition if Rush Order (6 Hrs)
				// //if and status is active
				// //set to late
				// if v.IsRushOrder != nil {
				// 	activationDateTimePlus6hrs := pointers.PrimativeToDateTime(*v.ActivationDateTime).Add(time.Hour * time.Duration(6))
				// 	if strings.ObjectTOString(v.Status) == constants.PipelineStatusActive && activationDateTimePlus6hrs.Before(time.Now().Local()) && *v.IsRushOrder {
				// 		newOrderAmount := *v.OrderFee - float64(*pipelineState.OPRush)
				// 		pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusLate, pointers.Float64(newOrderAmount)}

				// 	}
				// }

				//seting up cancelled after 2days of late
			}
			// else if strings.ObjectTOString(v.Status) == constants.PipelineStatusLate && duedatetime.Before(time.Now().Local()) {
			// 	pipelineStatuses[v.ID.Hex()] = NewPipeline{constants.PipelineStatusCancelled, nil}
			// }
		}
	}
	return pipelineStatuses

}
