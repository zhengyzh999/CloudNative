package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const README = "README.md"

// ReadLine1 一次性读取。适合小文件读取
func ReadLine1() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(bytes), "\n")
	for _, l := range split {
		fmt.Println(l)
	}
}

// ReadLine2 按行拆分并打印
// bufio通过对io模块的封装，提供了数据的缓冲功能，能一定程度上减少大数据块读写带来的开销
// 当发起读写操作时，会尝试从缓冲区读取数据，缓冲区没有数据后，才会从数据源获取。缓冲区默认大小为4Kb
func ReadLine2() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

// ReadLine3 通过scanner按行读取。单行默认大小64Kb
func ReadLine3() {
	file, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
