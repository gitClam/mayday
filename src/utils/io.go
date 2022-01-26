package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

var IO *io

type io struct{}

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

// Load 读取文件
func (io *io) Load(path string) ([]byte, *FilErr) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte(""), &FilErr{"File reading error"}
	}

	return data, nil
}

// Save 保存文件
func (io *io) Save(path string, data []byte) *FilErr {
	//待修改
	//如果文件a.txt已经存在那么会忽略权限参数，清空文件内容。文件不存在会创建文件赋予权限
	err := ioutil.WriteFile(path, data, 0777)
	if err != nil {
		return &FilErr{"File saveing error"}
	}
	return nil
}

// PathExists 判断文件是否存在
func (io *io) PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
