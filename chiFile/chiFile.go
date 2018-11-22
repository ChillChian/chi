package chiFile

import (
	"io/ioutil"
	"os"
)

// 写入文件

func WriteFile(path string, content []byte) (err error) {
	return ioutil.WriteFile(path, content, os.FileMode(0777))
}

func MustWriteFile(path string, content []byte) {
	if path == "" {
		panic(`[MustWriteFile] path == ""`)
		err := ioutil.WriteFile(path, content, os.FileMode(0777))
		if err != nil {
			panic(err)
		}
	}
}
