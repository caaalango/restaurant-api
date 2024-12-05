package dtohdls

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/calango-productions/api/internal/types"
	"github.com/gin-gonic/gin"
)

func GetUserData(ctx *gin.Context) (*types.UserData, error) {
	value, exists := ctx.Get("UserData")
	if !exists {
		return nil, errors.New("user data not found in context")
	}

	userData, ok := value.(*types.UserData)
	if !ok {
		return nil, errors.New("invalid user data format in context")
	}

	return userData, nil
}

func GetBody[T any](ctx *gin.Context) (*T, error) {
	var body T
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return nil, err
	}
	return &body, nil
}

func GetQueries[T any](ctx *gin.Context) (*T, error) {
	var result T
	resultValue := reflect.ValueOf(&result).Elem()
	resultType := resultValue.Type()

	missingParams := []string{}

	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		fieldName := field.Name

		queryValue := ctx.Query(fieldName)
		if queryValue == "" {
			missingParams = append(missingParams, fieldName)
			continue
		}

		fieldValue := resultValue.FieldByName(fieldName)
		if !fieldValue.IsValid() || !fieldValue.CanSet() || fieldValue.Kind() != reflect.String {
			return nil, fmt.Errorf("field '%s' is not valid, settable, or of type string", fieldName)
		}

		fieldValue.SetString(queryValue)
	}

	if len(missingParams) > 0 {
		return nil, fmt.Errorf("missing or empty query parameters: %v", missingParams)
	}

	return &result, nil
}

func GetParams[T any](ctx *gin.Context) (*T, error) {
	var result T
	resultValue := reflect.ValueOf(&result).Elem()
	resultType := resultValue.Type()

	missingParams := []string{}

	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		fieldName := field.Name

		paramValue := ctx.Param(fieldName)
		if paramValue == "" {
			missingParams = append(missingParams, fieldName)
			continue
		}

		fieldValue := resultValue.FieldByName(fieldName)
		if !fieldValue.IsValid() || !fieldValue.CanSet() || fieldValue.Kind() != reflect.String {
			return nil, fmt.Errorf("field '%s' is not valid, settable, or of type string", fieldName)
		}

		fieldValue.SetString(paramValue)
	}

	if len(missingParams) > 0 {
		return nil, fmt.Errorf("missing or empty URL parameters: %v", missingParams)
	}

	return &result, nil
}
