package gorm

func Query() {
	// 单条查询,第一条
	t := Teacher{}
	res := DB.First(&t)
	Println("单条查询,第一条", res.RowsAffected, res.Error, t)

	// 单条查询,最后一条
	t2 := Teacher{}
	res = DB.Last(&t)
	Println("单条查询,最后一条", res.RowsAffected, res.Error, t2)

	// 无排序,第一条
	t3 := Teacher{}
	res = DB.Take(&t)
	Println("无排序,第一条", res.RowsAffected, res.Error, t3)

	// 将结果填充到map，使用DB.Model()指定model时，此结构体中的特殊序标签用不了，如json序列化、数据类型不同等，需要除去这些字段
	result := map[string]interface{}{}
	res = DB.Model(&Teacher{}).Omit("Birthday", "Roles", "JobInfo2").First(&result)
	Println("map接收结果,Model填充", res.RowsAffected, res.Error, result)

	// 将结果填充到map，使用DB.Table()指定table时，无需往结构体中填充字段，不涉及类型匹配问题，所以不用Omit。但是First和Last用不了，只能使用Take
	result = map[string]interface{}{}
	res = DB.Table("teachers").Take(&result)
	Println("map接收结果,Table填充", res.RowsAffected, res.Error, result)

	var teachers []Teacher
	res = DB.Where("name=?", "nick").Or("name=?", "Joe").Order("id desc").Limit(10).Find(&teachers)
	Println("批量查询Find", res.RowsAffected, res.Error, teachers)
}
