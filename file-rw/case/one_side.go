package _case

import (
	"io"
	"log"
	"os"
	"path"
)

func OneSideReadWrite2Tar() {
	list := getFilesPath(sourceDir)
	for _, l := range list {
		_, name := path.Split(l)
		tarFileName := targetDir + "one-side/" + name
		// 文件写入
		OneSideReadWrite(l, tarFileName)
	}
}
func OneSideReadWrite(srcName, tarName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	tar, err := os.OpenFile(tarName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer tar.Close()
	buf := make([]byte, 512)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		tar.Write(buf[:n])
	}
}
