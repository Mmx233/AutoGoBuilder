package shell

import (
	"mmx/CallBack"
	"mmx/file"
	"os"
	"os/exec"
	"strings"
)

func Exec(command string)(string,error){
	cmd := exec.Command("/bin/bash", "-c",command)
	out,err:=cmd.CombinedOutput()
	return strings.Trim(string(out),"\n"),err
}

func Zip(path string,file string,target string)error{
	_,err:=Exec("cd "+path+" && zip -q "+target+" "+file)
	return err
}

func DoBuild(Name string,Dir string,Path string,GOOS string,GOARCH string){
	if _,err:=Exec("cd ./"+Dir+"/src/ && CGO_ENABLED=0 GOOS="+GOOS+" GOARCH="+GOARCH+" go build main.go");err != nil {
		CallBack.Error("未知错误")
		file.Remove(Dir,true)
		os.Exit(3)
	}
	out:=file.FindDist(Path)
	if out==""{
		CallBack.Error("编译 "+GOOS+" "+GOARCH+" 版失败")
		file.Remove(Dir,true)
		os.Exit(3)
	}
	if err:=Zip(Path,out,"../dist/"+Name+"_"+GOOS+"_"+GOARCH+".zip");err != nil {
		CallBack.Error("未知错误")
		file.Remove(Dir,true)
		os.Exit(3)
	}
	file.Remove(Path+"/"+out,false)
}
