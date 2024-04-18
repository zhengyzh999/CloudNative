package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

func FileReadCase() {
	open, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer func(open *os.File) {
		err = open.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("释放文件资源")
	}(open)
	buf := make([]byte, 1024)
	for {
		n, err2 := open.Read(buf)
		if err2 != nil && err2 != io.EOF {
			log.Fatal(err2)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[0:n])
	}
}
