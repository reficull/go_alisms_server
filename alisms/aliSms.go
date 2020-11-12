package alisms

import (
	"encoding/json"
	"fmt"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

//SmsInfo : info struct
type SmsInfo struct {
	AccessID    string
	AccessKEY   string
	SmsTemplate string
	SignName    string
	Cell        string
	Code        string
}

type RetMsg struct {
	Success int32  `json:"success"`
	Message string `json:"message"`
}

var msgMap map[string]string = GetResponsMsg()

//SendSMS send sms
func SendSMS(info SmsInfo) string {

	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(info.AccessID, info.AccessKEY) // dysms.New(ACCESSID, ACCESSKEY)

	// 短信发送
	//respSendSms, err := dysms.SendSms(uuid.New(), cell, "画学反应", "SMS_160145083", `{"code":"`+code+`"}`).DoActionWithException()
	respSendSms, err := dysms.SendSms(uuid.New(), info.Cell, info.SignName, info.SmsTemplate, `{"code":"`+info.Code+`"}`).DoActionWithException()
	retMsg := RetMsg{}

	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())

		if errMsg, ok := msgMap[(*respSendSms).GetCode()]; ok {
			retMsg.Message = errMsg
			retMsg.Success = 1
		} else {
			retMsg.Message = "短信发送失败"
			retMsg.Success = 0
		}
		//os.Exit(0)
		//ret = fmt.Sprintf("Fail,cell:%s,%s\n", info.Cell, respSendSms.Error())
	} else {
		retMsg.Message = *respSendSms.ErrorMessage.Message
		retMsg.Success = 1
		fmt.Printf("OK,cell:%s,code:%s,%s\n", info.Cell, info.Code, respSendSms.String())
	}
	body, e := json.Marshal(retMsg)
	if e != nil {
		fmt.Printf("json encode error:%s", e)
	}

	return string(body)

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

func GetResponsMsg() map[string]string {
	resMsg := make(map[string]string)
	resMsg["OK"] = "请求成功"
	resMsg["isp.RAM_PERMISSION_DENY"] = "RAM权限DENY"
	resMsg["isv.OUT_OF_SERVICE"] = "业务停机"
	resMsg["isv.PRODUCT_UN_SUBSCRIPT"] = "未开通云通信产品的阿里云客户"
	resMsg["isv.PRODUCT_UNSUBSCRIBE"] = "产品未开通"
	resMsg["isv.ACCOUNT_NOT_EXISTS"] = "账户不存在"
	resMsg["isv.ACCOUNT_ABNORMAL"] = "账户异常"
	resMsg["isv.SMS_TEMPLATE_ILLEGAL"] = "短信模板不合法"
	resMsg["isv.SMS_SIGNATURE_ILLEGAL"] = "短信签名不合法"
	resMsg["isv.INVALID_PARAMETERS"] = "参数异常"
	resMsg["isp.SYSTEM_ERROR"] = "系统错误"
	resMsg["isv.MOBILE_NUMBER_ILLEGAL"] = "非法手机号"
	resMsg["isv.MOBILE_COUNT_OVER_LIMIT"] = "手机号码数量超过限制"
	resMsg["isv.TEMPLATE_MISSING_PARAMETERS"] = "模板缺少变量"
	resMsg["isv.BUSINESS_LIMIT_CONTROL"] = "业务限流"
	resMsg["isv.INVALID_JSON_PARAM"] = "JSON参数不合法，只接受字符串值"
	resMsg["isv.BLACK_KEY_CONTROL_LIMIT"] = "黑名单管控"
	resMsg["isv.PARAM_LENGTH_LIMIT"] = "参数超出长度限制"
	resMsg["isv.PARAM_NOT_SUPPORT_URL"] = "不支持URL"
	resMsg["isv.AMOUNT_NOT_ENOUGH"] = "账户余额不足"

	return resMsg
}
