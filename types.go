package types

import (
	"bytes"
	"fmt"
	"reflect"
)

func Extract(targetType reflect.Type, items ...interface{}) (targetObj interface{}) {
	target := reflect.New(targetType)

	for _, item := range items {
		extract(&target, item)
	}

	target = reflect.Indirect(target)

	return target.Interface()
}

func extract(target *reflect.Value, item interface{}) {
	itemType := reflect.TypeOf(item)
	itemValue := reflect.ValueOf(item)

	targetValue := reflect.Indirect(*target)

	for i := 0; i < itemType.NumField(); i++ {
		itemTypeField := itemType.Field(i)
		itemValueField := itemValue.Field(i)

		targetField := targetValue.FieldByName(itemTypeField.Name)

		if !targetField.IsValid() {
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
