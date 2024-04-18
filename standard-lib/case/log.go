package _case

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stderr)
}
func LogCase() {
	var a, b = -1, -2
	_, err := sum(a, b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("a:%d,b:%d,两数求和出现错误: %s\n", a, b, err.Error())
	// Fatalf日志会退出应用程序
	//log.Fatalf("a:%d,b:%d,两数求和出现错误: %s\n", a, b, err.Error())
}
