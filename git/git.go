package git

import (
	"fmt"
	"math/rand"
	"mmx/CallBack"
	"mmx/file"
	"mmx/shell"
	"os"
	"time"
)

func Clone(url string,name string) string {
	rand.Seed(time.Now().UnixNano())
	temp:=fmt.Sprintf("%v%v",time.Now().Unix(),rand.Int())+name
	os.Mkdir(temp,700)
	os.Mkdir(temp+"/dist",700)
	os.Mkdir(temp+"/src",700)
	_,err:=shell.Exec("git clone "+url+" "+temp+"/src/")
	if err != nil {
		CallBack.Error("克隆仓库失败")
		file.Remove(temp,true)
		os.Exit(3)
	}
	return temp
}
