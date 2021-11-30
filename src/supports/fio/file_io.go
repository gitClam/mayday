package file_io

import (
	"io/ioutil"
)

type FIL_ERR struct {
	emsg string
}

func (fil_err FIL_ERR) Error() string {
	if fil_err.emsg == ""{
		return "其他未知错误"
	} else {
		return fil_err.emsg
	}
}

func Load(path string) ([]byte,*FIL_ERR) {
	
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte(""),&FIL_ERR{"File reading error"}
	}

	return data,nil
}
func Save(path string, data []byte) *FIL_ERR {
	//待修改
	//如果文件a.txt已经存在那么会忽略权限参数，清空文件内容。文件不存在会创建文件赋予权限
	err := ioutil.WriteFile(path, data, 0777) 
	if(err != nil){
		return 	&FIL_ERR{"File saveing error"}
	}
	return nil;
}
