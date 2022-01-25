package file_io

import (
	"io/ioutil"
)

type FilErr struct {
	msg string
}

func (filErr FilErr) Error() string {
	if filErr.msg == "" {
		return "其他未知错误"
	} else {
		return filErr.msg
	}
}

func Load(path string) ([]byte, *FilErr) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte(""), &FilErr{"File reading error"}
	}

	return data, nil
}
func Save(path string, data []byte) *FilErr {
	//待修改
	//如果文件a.txt已经存在那么会忽略权限参数，清空文件内容。文件不存在会创建文件赋予权限
	err := ioutil.WriteFile(path, data, 0777)
	if err != nil {
		return &FilErr{"File saveing error"}
	}
	return nil
}
