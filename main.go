package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"mmx/CallBack"
	"mmx/file"
	"mmx/git"
	"mmx/shell"
	"os"
	"time"
)

func main(){
	var repo string
	var backend string
	var name string
	var unix string
	flag.StringVar(&repo, "repo", "", "仓库url")
	flag.StringVar(&backend, "backend", "", "回调url")
	flag.StringVar(&name, "name", "", "文件名")
	flag.StringVar(&unix, "unix", fmt.Sprintf("%v",time.Now().Unix()), "随机字符串标识符")
	flag.Parse()

	dir:=git.Clone(repo,name,unix)
	CallBack.Receiver(backend,dir)
	path:=dir+"/src"
	defer file.Remove(path,true)

	if !file.Exists(path+"/main.go"){
		CallBack.Error("main.go不存在")
	}else{
		CallBack.Init()
		//编译
		defaultARCH:=[]string{"386","amd64"}
		sys := map[string][]string{
			"linux":   {"386", "amd64", "arm"},
			"darwin":  {"amd64"},
			"windows": defaultARCH,
		}
		for key,mips:=range sys{
			for _,mip:=range mips {
				shell.DoBuild(name, dir, path, key, mip)
			}
		}
		if err:=shell.Zip(dir+"/dist/*",dir+"/dist/"+name+"_all.zip");err!=nil{
			CallBack.Error("压缩过程出错")
		}
	}
	//对接fileBed
	files,_:=ioutil.ReadDir(dir+"/dist/")
	os.Mkdir("FILE/"+dir,700)
	for _,file:=range files{
		os.Rename(dir+"/dist/"+file.Name(),"FILE/"+dir+"/"+file.Name())
	}
	os.RemoveAll(dir)
	data:=map[string]interface{}{
		"SignKey":dir,
	}
	CallBack.Success(data)
}