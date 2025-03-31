package godb

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Statement interface {
	buildQuery() string
	DB() *sql.DB
}

func Execute[T any](stmt Statement, record T) ([]T, error) {
	structType := reflect.TypeOf(record)

	if structType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("record must be of type struct")
	}

	queryStr := stmt.buildQuery()

	rows, err := stmt.DB().Query(queryStr)

	if err != nil {
		log.Printf("failed to fetch rows: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	var results [][]interface{}
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		log.Printf("error gettings column types: %s\n", err)
		return nil, err
	}

	for rows.Next() {
		values := make([]interface{}, len(columnTypes))
		scanArgs := make([]interface{}, len(columnTypes))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			log.Printf("error scanning row: %s\n", err)
			return nil, err
		}

		results = append(results, values)
	}

	err = rows.Err()

	if err != nil {
		log.Printf("error iterating rows: %s\n", err)
		return nil, err
	}

	slice := reflectStruct(structType, results)

	return slice.Interface().([]T), nil

}

func reflectStruct(structType reflect.Type, values [][]interface{}) reflect.Value {
	count := len(values)
	sliceType := reflect.SliceOf(structType)
	slice := reflect.MakeSlice(sliceType, count, count)

	for i, row := range values {
		element := reflect.New(structType).Elem()

		for j, val := range row {
			field := structType.Field(j)
			fieldValue := element.Field(j)

			if fieldValue.Kind() != reflect.Invalid {
				switch field.Type.Kind() {
				case reflect.Int:
					if v, ok := val.(int64); ok {
						fieldValue.SetInt(v)
					}
				case reflect.String:
					if v, ok := val.(string); ok {
						fieldValue.SetString(v)
					}
				case reflect.Bool:
					if v, ok := val.(bool); ok {
						fieldValue.SetBool(v)
					}
				}

			}
		}
		slice.Index(i).Set(element)
	}

	return slice
}
