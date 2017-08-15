package types

import (
	"bytes"
	"fmt"
	"reflect"
)

func getPointerType(tp reflect.Type) reflect.Type {
	for tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	return tp
}

func getPointerValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	return val
}

func Extract(target interface{}, items ...interface{}) (targetObj interface{}) {
	switch tp := target.(type) {
	case reflect.Type:
		targetObj = extractNewInstance(tp, items)
	default:
		// targetType := getPointerType(reflect.TypeOf(target))
		targetValue := getPointerValue(reflect.ValueOf(target))
		if !targetValue.CanAddr() {
			targetObj = extractNewInstance(
				getPointerType(reflect.TypeOf(target)),
				items,
			)
			return
		}
		for _, item := range items {
			extract(&targetValue, item)
		}
		targetObj = targetValue.Interface()
	}
	return
}

func extractNewInstance(targetType reflect.Type, items []interface{}) (targetObj interface{}) {
	target := reflect.New(targetType)

	for _, item := range items {
		extract(&target, item)
	}

	target = reflect.Indirect(target)

	return target.Interface()
}

func extract(target *reflect.Value, item interface{}) {
	itemType := getPointerType(reflect.TypeOf(item))
	itemValue := getPointerValue(reflect.ValueOf(item))
	targetValue := getPointerValue(*target)

	for i := 0; i < itemType.NumField(); i++ {
		itemTypeField := itemType.Field(i)
		itemValueField := itemValue.Field(i)

		targetField := targetValue.FieldByName(itemTypeField.Name)

		if !targetField.IsValid() || !targetField.CanSet() {
			continue
		}
		targetField.Set(itemValueField)

	}
}

func Println(items ...interface{}) {
	for _, item := range items {
		println(item)
	}
}

func println(item interface{}) {
	buffer := bytes.Buffer{}
	itemType := reflect.TypeOf(item)

	itemValue := reflect.ValueOf(item)

	if itemType.Kind() != reflect.Struct {
		fmt.Println(itemType)
		return
	}

	buffer.WriteString(
		fmt.Sprintf("%s : { \n", itemType),
	)
	for i := 0; i < itemType.NumField(); i++ {
		itemTypeField := itemType.Field(i)
		itemValueField := itemValue.Field(i)
		buffer.WriteString(
			fmt.Sprintf("\t\"%s\" : %s ", itemTypeField.Name, parseValue2String(itemValueField)),
		)
		if i != itemType.NumField()-1 {
			buffer.WriteString(",\n")
		}
	}
	buffer.WriteString("\n}")
	fmt.Println(buffer.String())
}

func PrintlnInOneLine(items ...interface{}) {
	for _, item := range items {
		printlnInOneLine(item)
	}
}

func printlnInOneLine(item interface{}) {
	buffer := bytes.Buffer{}
	itemType := reflect.TypeOf(item)

	itemValue := reflect.ValueOf(item)

	if itemType.Kind() != reflect.Struct {
		fmt.Println(itemType)
		return
	}

	buffer.WriteString(
		fmt.Sprintf("%s : { ", itemType),
	)
	for i := 0; i < itemType.NumField(); i++ {
		itemTypeField := itemType.Field(i)
		itemValueField := itemValue.Field(i)
		buffer.WriteString(
			fmt.Sprintf("\"%s\" : %s ", itemTypeField.Name, parseValue2String(itemValueField)),
		)
		if i != itemType.NumField()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	fmt.Println(buffer.String())
}

func parseValue2String(value reflect.Value) string {
	if value.Kind() == reflect.String {
		return fmt.Sprintf("\"%s\"", value)
	}
	return fmt.Sprintf("%v", value)
}
