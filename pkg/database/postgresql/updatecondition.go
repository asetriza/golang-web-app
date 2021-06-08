package postgresql

import (
	"fmt"
	"reflect"
	"strings"
)

// UpdateConditionFromStruct
// Creates condition string (valueOfTag = :valueOfTag) for updating record in database.
func UpdateConditionFromStruct(in interface{}) string {
	var updateConditionString string
	valuesOfStruct := reflect.ValueOf(in).Elem()
	typeOfStruct := reflect.TypeOf(in).Elem()
	for i := 0; i < valuesOfStruct.NumField(); i++ {
		if reflect.ValueOf(in).Elem().Field(i).String() != "" {
			if fmt.Sprintf("%v", reflect.ValueOf(in).Elem().Field(i)) != "0" {
				t := reflect.TypeOf(in).Elem()
				fieldOfStruct, _ := t.FieldByName(typeOfStruct.Field(i).Name)
				valueOfTag, _ := fieldOfStruct.Tag.Lookup("db")
				if reflect.ValueOf(in).Elem().Field(i).String() == "delete" {
					updateConditionString += valueOfTag + " = null, "
				} else {
					updateConditionString += valueOfTag + " = :" + valueOfTag + ", "
				}
			}
		}
	}
	updateConditionString = strings.TrimSuffix(updateConditionString, ", ")
	return updateConditionString
}
