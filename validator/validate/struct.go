package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name            string         `v:"required,alphaunicode"`
	Age             uint8          `v:"gte=10,lte=30"`
	Phone           string         `v:"required,e164"`
	Email           string         `v:"required,email"`
	FavouriteColor1 string         `v:"iscolor"`
	FavouriteColor2 string         `v:"hexcolor|rgb|rgba|hsl|hsla"`
	Address         *Address       `v:"required"`
	ContactUser     []*ContactUser `v:"required,gte=1,dive"` // 深入下一层进行验证
	Hobby           []string       `v:"required,gte=2,dive,required,gte=2,alphaunicode"`
}
type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}
type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=20,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Email Phone"`
}

func StructValidate() {
	v := validate
	address := &Address{
		Province: "湖南",
		City:     "长沙",
	}
	contactUser1 := &ContactUser{
		Name:  "张三",
		Age:   30,
		Phone: "+8612345678901",
		//Email:   "+8612345678901@qq.com",
		//Address: address,
	}
	contactUser2 := &ContactUser{
		Name:    "李四",
		Age:     47,
		Phone:   "+8612345678902",
		Email:   "+8612345678902@qq.com",
		Address: address,
	}
	user := &User{
		Name:            "nick",
		Age:             18,
		Phone:           "+86138001380000",
		Email:           "nick@gmail.com",
		FavouriteColor1: "#ffffff",
		FavouriteColor2: "rgb(255,255,255)",
		Address:         address,
		ContactUser:     []*ContactUser{contactUser1, contactUser2},
		Hobby:           []string{"乒乓球", "羽毛球"},
	}
	err := v.Struct(user)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
