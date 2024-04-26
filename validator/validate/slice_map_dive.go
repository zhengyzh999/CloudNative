package validate

type SliceStruct struct {
	OpCode int    `v:"eq=1|eq=2"`
	Op     string `v:"required"`
}

func SliceValidate() {
	v := validate
	var err error
	slice1 := []string{"12345", "67890", "12345678901"}
	err = v.Var(slice1, "gte=3,dive,required,gte=5,lte=10")
	outResult("slice1", &err)

	slice2 := [][]string{
		{"12345", "67890"},
		{"12345", "67890", "12345678901"},
		{"12345", "67890", "12345678901"},
	}
	err = v.Var(slice2, "gte=3,dive,gte=3,dive,required,gte=5,lte=10,number")
	outResult("slice2", &err)

	slice3 := []*SliceStruct{
		{OpCode: 1,
			Op: "切片操作",
		},
		{OpCode: 2,
			Op: "切片操作",
		},
		{OpCode: 3,
			Op: "切片操作",
		},
	}
	err = v.Var(slice3, "gte=2,dive")
	outResult("slice3", &err)
}

func MapValidate() {
	v := validate
	var err error
	mp1 := map[string]string{
		"AA": "12345",
		"B":  "12345998765",
		"C":  "12345",
	}
	err = v.Var(mp1, "gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outResult("map1", &err)

	mp2 := map[string]map[string]string{
		"A": {
			"Ax": "12345",
			"Ay": "1234567890",
			"Az": "12345",
		},
		"B": {
			"Bx": "12345",
			"By": "1234567890",
			"Bz": "12345",
		},
		"C": {
			"Cx": "12345",
			"Cy": "1234567890",
			"Cz": "12345",
		},
	}
	err = v.Var(mp2, "gte=2,dive,keys,len=1,alpha,endkeys,required,gte=3,dive,keys,len=2,alpha,endkeys,required,gte=5,lte=10,number")
	outResult("map2", &err)
}
