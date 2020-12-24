package file

import "os"

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Remove(path string,Isdir bool){
	if Isdir{
		_=os.RemoveAll(path)
	}else {
		_ = os.Remove(path)
	}
}

func FindDist(Path string)string{
	if Exists(Path+"/main"){
		return "main"
	}
	if Exists(Path+"/main.exe"){
		return "main.exe"
	}
	return ""
}
