package CallBack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var backend string

func makeCall(data interface{}){
	body,_:=json.Marshal(data)
	_, _ = http.Post(backend, "application/json", bytes.NewReader(body))
}

func Receiver(url string){
	backend=url
}

func Error(msg string){
	type errorCall struct {
		Status string `json:"status"`
		Msg string `json:"msg"`
	}
	var temp errorCall
	temp.Status="error"
	temp.Msg=msg
	makeCall(temp)
}

func Success(){

}

