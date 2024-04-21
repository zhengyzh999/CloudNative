package _case

import (
	"io"
	"log"
	"os"
	"path"
)

func CopyDir2Dir() {
	list := getFilesPath(sourceDir)
	for _, f := range list {
		_, name := path.Split(f)
		targetFileName := targetDir + "copy/" + name
		copyFile(f, targetFileName)

	}
}

// 复制文件
func copyFile(srcName, tarName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	tar, err := os.OpenFile(tarName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer tar.Close()
	return io.Copy(tar, src)
}
