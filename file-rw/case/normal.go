package _case

import (
	"log"
	"os"
	"path"
)

func ReadWriteFiles() {
	list := getFilesPath(sourceDir)
	for _, f := range list {
		bytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		_, name := path.Split(f)
		tarName := targetDir + "normal/" + name
		err = os.WriteFile(tarName, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
