package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

var IO *Io

type Io struct{}

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
func (i *Io) Load(path string) ([]byte, *FilErr) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte(""), &FilErr{"File reading error"}
	}

	return data, nil
}

// Save 保存文件（没有就创建，删除并覆盖）
func (i *Io) Save(path string, file multipart.File) (err1 error) {

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		err1 = err
		return
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			err1 = err
			return
		}
	}(out)

	_, err = io.Copy(out, file)
	if err != nil {
		err1 = err
		return
	}

	return nil
}

// PathExists 判断文件是否存在
func (i *Io) PathExists(path string) (bool, error) {
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
