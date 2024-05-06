package gorm

import (
	"log"
	"time"
)

var teacherTemp = Teacher{
	Name:         "nick",
	Age:          40,
	WorkingYears: 10,
	Email:        "nick@163.com",
	Birthday:     time.Now().Unix(),
	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},
	Roles: []string{"普通用户", "讲师"},
	JobInfo: Job{
		Title:    "金牌讲师",
		Location: "北京北京",
	},
	JobInfo2: Job{
		Title:    "金牌讲师",
		Location: "北京北京",
	},
}

func CreateRecord() {
	// 普通插入
	t := teacherTemp
	res := DB.Create(&t)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	Println(res.RowsAffected, res.Error, t)

	// 正向选择
	t1 := teacherTemp
	res = DB.Select("Name", "Age").Create(&t1)
	Println(res.RowsAffected, res.Error, t1)

	// 反向选择
	t2 := teacherTemp
	res = DB.Omit("Email", "Birthday").Create(&t2)
	Println("t2", res.RowsAffected, res.Error, t2)

	// 批量操作
	var teachers = []Teacher{{Name: "Jack", Age: 42}, {Name: "Lucy", Age: 34}, {Name: "Joe", Age: 83}}
	DB.Create(teachers)
	for _, t := range teachers {
		Println("批量操作", t.ID)
	}

	// 批量插入，设定批处理大小
	var teachers1 = []Teacher{{Name: "Jack", Age: 42}, {Name: "Lucy", Age: 34}, {Name: "Joe", Age: 83}}
	DB.CreateInBatches(teachers1, 2)
	for _, t := range teachers1 {
		Println("批量操作,设置批处理大小", t.ID)
	}
}
