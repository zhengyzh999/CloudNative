package _case

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		Id    int64
		Name  string
		Hobby []string
	}
	type outUser struct {
		Id    int64
		Name  string
		Hobby []string
	}
	u := user{
		Id:    1,
		Name:  "tom",
		Hobby: []string{"唱", "跳"},
	}
	outU := outUser{}
	res := copy(&outU, u)
	fmt.Println("res = ", res, outU)

	listUser := []user{
		{
			Id:    1,
			Name:  "tom",
			Hobby: []string{"唱", "跳"},
		},
		{
			Id:    2,
			Name:  "tom1",
			Hobby: []string{"唱1", "跳1"},
		},
		{
			Id:    3,
			Name:  "tom2",
			Hobby: []string{"唱1", "跳2"},
		},
	}
	list := sliceColumn(listUser, "Hobby")
	fmt.Println("list = ", list)
}

func sliceColumn(slice interface{}, column string) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		val := v.FieldByName(column)
		return val.Interface()
	}
	if v.Kind() != reflect.Slice {
		return nil
	}
	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, _ := t.FieldByName(column)
	sliceType := reflect.SliceOf(f.Type)
	s := reflect.MakeSlice(sliceType, 0, 0)
	for i := 0; i < v.Len(); i++ {
		o := v.Index(i)
		if o.Kind() == reflect.Struct {
			val := o.FieldByName(column)
			s = reflect.Append(s, val)
		}
		if o.Kind() == reflect.Ptr {
			v1 := o.Elem()
			val := v1.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}
	return s.Interface()
}

func copy(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	// 如果是指针类型，则获取其值
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须为指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()
	if sValue.Kind() != reflect.Struct {
		return errors.New("源对象必须为struct或struct的指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目标对象必须为struct的指针")
	}
	destObj := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if destField.Type != sourceField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObj.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObj.Elem())
	return nil
}
