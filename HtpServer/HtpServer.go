package  HtpServer

import(
	"fmt"
	"log"
	"net/http"
    "AliSMS"
)

type CommandType int

const (
	//GetCommand ...
	GetCommand = iota
	//SetCommand ...
	SetCommand
	//IncCommand ...
	IncCommand
	//SMSCommand ...
	SMSCommand
)
type Command struct {
	ty        CommandType
	code      string
	cell      string
	replyChan chan string
}
type Server struct {
	Cmds chan<- Command
}

func StartProcessManager(initvals map[string]float32, info AliSMS.SmsInfo) chan<- Command {
	counters := make(map[string]float32)

	for k, v := range initvals {
		counters[k] = v
	}
	cmds := make(chan Command)
	go func() {
		for cmd := range cmds {
			switch cmd.ty {

			case SMSCommand:
				var ret string
				fmt.Println("send code:", cmd.code, " cell:", cmd.cell)
				fmt.Printf("%v", info)
				info.Code = cmd.code
				info.Cell = cmd.cell
				ret = AliSMS.SendSMS(info)
				//fmt.Printf("ct command logic  s1:%s,s2:%s\n ret:%s",cmd.str1,cmd.str2,ret)

				cmd.replyChan <- ret
			default:
				//log.Fatal("unknown command type", cmd.ty)
			}
		}
	}()
	return cmds
}

//Get method
func (s *Server) Get(w http.ResponseWriter, req *http.Request) {
	log.Printf("get %v", req)
	name := req.URL.Query().Get("name")
	replyChan := make(chan string)
	s.Cmds <- Command{ty: GetCommand, replyChan: replyChan}
	reply := <-replyChan
	fmt.Fprintf(w, "%s: %s\n", name, reply)
}

//Sms handle sms call
func (s *Server) Sms(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	cell := req.URL.Query().Get("cell")
	replyChan := make(chan string)
	s.Cmds <- Command{ty: SMSCommand, code: code, cell: cell, replyChan: replyChan}
	//TODO:check info

	reply := <-replyChan
	fmt.Fprint(w, reply)
}

