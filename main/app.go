package main

import (
	"fmt"
	"log"
	"net/http"

	alisms "github.com/reficull/go_alisms_server/alisms"
	htpserver "github.com/reficull/go_alisms_server/htpserver"
)

func main() {
	fmt.Printf("hello")

	info := alisms.SmsInfo{
		AccessID:    "LTAI2YjtA8kIpW6k",
		AccessKEY:   "WWhcX9jcKXsGVkSo8RPQcIQdsaerz3",
		SmsTemplate: "SMS_160145083",
		SignName:    "画学反应",
		Cell:        "",
		Code:        "",
	}

	server := htpserver.Server{Cmds: htpserver.StartProcessManager(map[string]float32{"i": 0, "j": 0}, info)}

	http.HandleFunc("/get", server.Get)

	http.HandleFunc("/sms", server.Sms)

	log.Printf("compare text service Going to listen on port 8881\n")
	log.Fatal(http.ListenAndServe(":8881", nil))

}
