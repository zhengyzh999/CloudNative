package util

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)
const defaultNum = 100

func ParseSize(size string) (int64, string) {
	//time.Sleep(time.Nanosecond * 500)
	rep, _ := regexp.Compile("[0-9]+")
	unit := string(rep.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}
	if num == 0 {
		//log.Println("ParseSize 仅支持B、KB、MB、GB、TB、PB")
		num = defaultNum
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizeStr
}
