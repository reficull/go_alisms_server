module sms

replace HtpServer => ./HtpServer

replace AliSMS => ./AliSMS

go 1.15

require (
	AliSMS v0.0.0-00010101000000-000000000000 // indirect
	HtpServer v0.0.0-00010101000000-000000000000 // indirect
	github.com/GiterLab/aliyun-sms-go-sdk v0.0.0-20180108012719-fcc9f11de968 // indirect
	github.com/GiterLab/urllib v0.0.0-20200820124023-82571a63c776 // indirect
	github.com/tobyzxj/uuid v0.0.0-20140223123307-aa0153c14395 // indirect
)
