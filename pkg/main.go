package generalutils

import (
	"errors"
	"fmt"
	"reflect"
)

func BuildEntityToPersist[T any](params interface{}) (*T, error) {
	var entity T

	entityType := reflect.TypeOf(entity)
	entityPtr := reflect.New(entityType)

	entityValue := entityPtr.Interface()

	err := mapStructFields(params, entityValue)
	if err != nil {
		return nil, fmt.Errorf("failed to map fields: %v", err)
	}

	return entityValue.(*T), nil
}

func mapStructFields(input interface{}, output interface{}) error {
	inputVal := reflect.ValueOf(input)
	outputVal := reflect.ValueOf(output)

	if inputVal.Kind() != reflect.Ptr || outputVal.Kind() != reflect.Ptr {
		return errors.New("both input and output must be pointers")
	}

	inputVal = inputVal.Elem()
	outputVal = outputVal.Elem()

	if inputVal.Kind() != reflect.Struct || outputVal.Kind() != reflect.Struct {
		return errors.New("both input and output must be structs")
	}

	inputType := inputVal.Type()

	for i := 0; i < inputVal.NumField(); i++ {
		fieldName := inputType.Field(i).Name
		outputField := outputVal.FieldByName(fieldName)

		if outputField.IsValid() && outputField.CanSet() {
			outputField.Set(inputVal.Field(i))
		}
	}

	return nil
}
