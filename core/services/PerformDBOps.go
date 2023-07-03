package services

import (
	"reflect"
	"strings"
	"xorm.io/xorm"
)

// GetPrimaryKeys is a generic function that takes a struct and returns a slice of its primary key values.
// It uses reflection to iterate over the fields of the struct and checks the 'xorm' tag to identify primary keys.
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

// UpsertDB is a generic function that tries to insert a new row into the database.
// If the insert fails (usually due to a primary key conflict), it updates the existing row.
func UpsertDB[T any](data T, engine *xorm.Engine) error {
	pks := GetPrimaryKeys(data)

	_, err := engine.Insert(&data)
	if err != nil {
		_, err = engine.ID(pks).Update(&data)
	}

	return err
}

// InsertDB is a generic function that inserts a new row into the database.
func InsertDB[T any](data T, engine *xorm.Engine) error {
	_, err := engine.Insert(&data)

	return err
}

// UpdateDB is a generic function that updates an existing row in the database.
// It uses the primary key values of the struct to identify the row to update.
func UpdateDB[T any](data T, engine *xorm.Engine) error {
	_, err := engine.ID(GetPrimaryKeys(data)).Update(&data)
	return err
}

// QueryDB is a generic function that queries the database and fills the passed struct with the result.
// The condition parameter is a SQL WHERE clause.
// Note: The 'data' parameter should be a pointer to an array type.
func QueryDB[T any](condition string, data *T, engine *xorm.Engine) error {
	err := engine.Where(condition).Find(data)
	return err
}
