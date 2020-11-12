package htpserver

import (
	"fmt"
	"log"
	"net/http"

	//    "os"

	"github.com/reficull/go_alisms_server/aliSms"
)

//CommandType command type
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

// Command struct
type Command struct {
	ty        CommandType
	code      string
	cell      string
	replyChan chan string
}

//Server struct
type Server struct {
	Cmds chan<- Command
}

//StartProcessManager initializer
func StartProcessManager(initvals map[string]float32, info aliSms.SmsInfo) chan<- Command {
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
				ret = aliSms.SendSMS(info)
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

/*
func main() {
    server := Server{startProcessManager(map[string]int{"i":0,"j":0})}
    http.HandleFunc("/get", server.get)
    http.HandleFunc("/set", server.set)
    http.HandleFunc("/inc", server.inc)
    http.HandleFunc("/ct", server.ct)

    portnum := 8000
    if len(os.Args) > 1 {
        portnum, _ = strconv.Atoi(os.Args[1])

    }
    log.Printf("Going to listen on port %d\n", portnum)
    log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))

}
*/
