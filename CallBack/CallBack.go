package CallBack

import (
	"bytes"
	"encoding/json"
	"mmx/file"
	"net/http"
	"os"
)

var backend string
var key string

func makeCall(data interface{}){
	body,_:=json.Marshal(data)
	_, _ = http.Post(backend, "application/json", bytes.NewReader(body))
}

func Receiver(url string,I string){
	backend=url
	key=I
}

func Error(msg string){
	type errorCall struct {
		Status string `json:"status"`
		Signkey string `json:"signkey"`
		Msg string `json:"msg"`
	}
	var temp errorCall
	temp.Status="error"
	temp.Signkey=key
	temp.Msg=msg
	makeCall(temp)
	file.Remove(key,true)
	os.Exit(3)
}

func Success(data interface{}){
	type successCall struct {
		Status string `json:"status"`
		Signkey string `json:"signkey"`
		Data interface{} `json:"data"`
	}
	var temp successCall
	temp.Status="success"
	temp.Signkey=key
	temp.Data=data
	makeCall(temp)
}

func Init(){
	type InitCall struct {
		Status string `json:"status"`
		SignKey string `json:"sign_key"`
	}
	var temp InitCall
	temp.Status="init"
	temp.SignKey=key
	makeCall(temp)
}

