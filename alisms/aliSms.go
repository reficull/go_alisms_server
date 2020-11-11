package aliSms

import (
	"fmt"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// modify it to yours
const (
	ACCESSID  = ""
	ACCESSKEY = ""
)

func SendSMS(string code, string cell) string {
	var ret string
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(ACCESSID, ACCESSKEY) // dysms.New(ACCESSID, ACCESSKEY)

	// 短信发送
	respSendSms, err := dysms.SendSms(uuid.New(), cell, "画学反应", "SMS_160145083", `{"code":"`+code+`"}`).DoActionWithException()
	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())
		//os.Exit(0)
		ret = fmt.Printf("Fail,cell:%s,%s\n", cell, respSendSms.Error())
	} else {
		ret = fmt.Printf("OK,cell:%s,code:%s,%s\n", cell, code, respSendSms.String())
	}
	return ret

	// 查询短信
	/*
		respQuerySendDetails, err := dysms.QuerySendDetails("612710515335092485^0", "1375821****", "10", "1", "20180107").DoActionWithException()
		if err != nil {
			fmt.Println("query sms failed", err, respQuerySendDetails.Error())
			os.Exit(0)
		}
		fmt.Println("query sms succeed", respQuerySendDetails.String())
	*/
}
