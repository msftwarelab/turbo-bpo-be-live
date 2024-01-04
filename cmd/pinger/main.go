package main

import (
	//"context"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/slack-go/slack"
)

func main() {

	if len(os.Args) != 2 {
		log.Error("%s", "parameter required: [development,staging,production]")
		os.Exit(1)
	}

	config.Init("pinger", os.Args[1])

	log.Init("pinger-job")
	log.Info("Started pinger job")
	//run jobs every 60 sec

	frequency := time.Second * time.Duration(config.AppConfig.GetInt64("runEveryInSec"))
	//ctx, _ := context.WithCancel(context.Background())

	for {

		resp, err := http.Get(config.AppConfig.GetString("graphQLSserver"))
		if err != nil {
			log.Error(" http err %v : ", err)
		}

		// Print the HTTP Status Code and Status Name
		//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		// log.Info("sending slack notifiction...")
		// slackNotification(config.AppConfig.GetString("slackWebHookURL"), config.AppConfig.GetString("graphQLSserver"))
		// log.Info("done...")
		if resp != nil {
			if resp.StatusCode != 200 {

				log.Error("server error: status down...", nil)
				log.Info("restarting graphql server...")

				cmd := exec.Command("systemctl", "restart", "turbo-be-staging.service")
				if err := cmd.Run(); err != nil {
					log.Error(" cmd err  %v : ", err)
				} else {
					log.Info("done......")
					log.Info("sending slack notification...")
					slackNotification(config.AppConfig.GetString("slackWebHookURL"), config.AppConfig.GetString("graphQLSserver"))
					log.Info("done...")
				}
			}

		} else {

			log.Error("server error: status down...", nil)
			log.Info("restarting graphql server...")

			cmd := exec.Command("systemctl", "restart", "turbo-be-staging.service")
			if err := cmd.Run(); err != nil {
				log.Error(" cmd err : %v", err)
			} else {
				log.Info("done......")
				log.Info("sending slack notification...")
				slackNotification(config.AppConfig.GetString("slackWebHookURL"), config.AppConfig.GetString("graphQLSserver"))
				log.Info("done...")
			}

		}
		time.Sleep(frequency)

	}
}

func slackNotification(slackWebHookURL, serverUrl string) {
	text := fmt.Sprintf("<!channel> this server: %s  is detected down, we applied self heal.", serverUrl)
	attachment := slack.Attachment{
		Color:         "red",
		Fallback:      "You successfully posted by Incoming Webhook URL!",
		AuthorName:    "Turbo BPO Pinger",
		AuthorSubname: "github.com",
		AuthorLink:    "https://github.com/slack-go/slack",
		AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
		Text:          text,
		Footer:        "bevs bot",
		FooterIcon:    "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(slackWebHookURL, &msg)
	if err != nil {
		log.Error("slack notification error : %v", err)
	}
}
