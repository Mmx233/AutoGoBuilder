package git

import (
	"mmx/CallBack"
	"mmx/file"
	"mmx/shell"
	"os"
)

func Clone(url string,name string,unix string) string {
	temp:=unix+name
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
