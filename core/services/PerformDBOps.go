package services

import (
	"fmt"
	"reflect"
	"strings"
	"xorm.io/xorm"
)

func GetPrimaryKeys[T any](data T) []interface{} {
	v := reflect.ValueOf(data)
	var keys []interface{}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		if strings.Contains(tag.Get("xorm"), "pk") {
			keys = append(keys, v.Field(i).Interface())
		}
	}

	return keys
}

func Upsert[T any](data T, engine *xorm.Engine) error {
	pks := GetPrimaryKeys(data)

	_, err := engine.Insert(&data)
	if err != nil {
		fmt.Println(err)
		_, err = engine.ID(pks).Update(&data)
	}

	return err
}

func Insert[T any](data T, engine *xorm.Engine) error {
	_, err := engine.Insert(&data)

	return err
}

func Update[T any](data T, engine *xorm.Engine) error {
	_, err := engine.ID(GetPrimaryKeys(data)).Update(&data)
	return err
}
