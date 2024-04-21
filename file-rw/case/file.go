package _case

import (
	"log"
	"os"
	"strings"
)

// 源文件目录
const sourceDir = "source-file/"

// 目标文件目录
const targetDir = "target-file/"

func getFilesPath(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fullName := strings.Trim(dir, "/") + "/" + f.Name()
		list = append(list, fullName)
	}
	return list
}
