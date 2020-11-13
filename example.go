package main

import(
	"fmt"
	"log"
	"net/http"
    "HtpServer"
    "AliSMS"
)

func main(){
	fmt.Printf("hello")

	info := AliSMS.SmsInfo{
		AccessID:    "your access id",
		AccessKEY:   "you access secret",
		SmsTemplate: "SMS_160145083",
		SignName:    "your sign name",
		Cell:        "",
		Code:        "",
	}

	server := HtpServer.Server{Cmds: HtpServer.StartProcessManager(map[string]float32{"i": 0, "j": 0}, info)}

	http.HandleFunc("/get", server.Get)

	http.HandleFunc("/sms", server.Sms)

	log.Printf("compare text service Going to listen on port 8881\n")
	log.Fatal(http.ListenAndServe(":8881", nil))
}
