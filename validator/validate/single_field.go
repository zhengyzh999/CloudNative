package validate

import "time"

// SingleFieldValidate 单个字段格式验证
func SingleFieldValidate() {
	v := validate
	var err error
	// 布尔
	var b bool
	err = v.Var(b, "boolean")
	outResult("boolean", &err)

	// 数字
	var i = "100"
	err = v.Var(i, "number")
	outResult("number", &err)

	var f = 100.23
	err = v.Var(f, "numeric")
	outResult("numeric", &err)
	err = v.Var(f, "number")
	outResult("number", &err)

	// 字符串
	var str = "abcdefg"
	err = v.Var(str, "alpha")
	outResult("alpha", &err)

	// 切片
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	err = v.Var(slice, "max=15,min=2")
	outResult("slice", &err)

	// 集合
	var mp = make(map[int]int)
	mp[1] = 1
	mp[2] = 2
	mp[3] = 3
	err = v.Var(mp, "max=15,min=2")
	outResult("map", &err)

	// 时间
	var timeStr = time.Now().Format("2006-01-02 15:04:05")
	err = v.Var(timeStr, "datetime=2006-01-02 15:04:05")
	outResult("datetime", &err)

	// 对比字段是否相等
	s1 := "abc"
	s2 := "abc"
	err = v.VarWithValue(s1, s2, "eqfield")
	outResult("eqfield", &err)

	i1 := 10
	i2 := 20
	err = v.VarWithValue(i1, i2, "ltfield")
	outResult("ltfield", &err)
}
