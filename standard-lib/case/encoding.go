package _case

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// EncodingCase encoding包使用，编码及解码
func EncodingCase() {
	type user struct {
		Id   int64
		Name string
		Age  uint8
	}
	u := user{
		Id:   12,
		Name: "zyz",
		Age:  33,
	}
	// json序列化和反序列化
	marshal, err := json.Marshal(u)
	fmt.Println(marshal, err)
	u1 := user{}
	err = json.Unmarshal(marshal, &u1)
	fmt.Println(u1, err)

	// base64编解码
	str := base64.StdEncoding.EncodeToString(marshal)
	fmt.Println("str = ", str)
	decodeString, err := base64.StdEncoding.DecodeString(str)
	fmt.Println("decodeString = ", decodeString, err)

	// 十六进制编解码
	str1 := hex.EncodeToString(marshal)
	fmt.Println("str1 = ", str1)
	hexDecodeString, err := hex.DecodeString(str1)
	fmt.Println("hexDecodeString = ", hexDecodeString, err)

}
