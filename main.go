package main

import (
	"flag"
	"mmx/CallBack"
	"mmx/file"
	"mmx/git"
	"mmx/shell"
)

func main(){

	var repo string
	var backend string
	var name string
	flag.StringVar(&repo, "repo", "", "仓库url")
	flag.StringVar(&backend, "backend", "", "回调url")
	flag.StringVar(&name, "name", "", "文件名")
	flag.Parse()
	CallBack.Receiver(backend)

	dir:=git.Clone(repo,name)
	path:=dir+"/src"
	defer file.Remove(path,true)

	if !file.Exists(path+"/main.go"){
		CallBack.Error("不存在")
	}else{
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
}