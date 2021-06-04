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
		f1 := valuesOfStruct.Field(i)
		if f1.IsZero() == false {
			t := reflect.TypeOf(in).Elem()
			fieldOfStruct, _ := t.FieldByName(typeOfStruct.Field(i).Name)
			valueOfTag, _ := fieldOfStruct.Tag.Lookup("db")
			if fmt.Sprintf("%+v", f1.Elem()) == "delete" {
				updateConditionString += valueOfTag + " = null, "
			} else {
				updateConditionString += valueOfTag + " = :" + valueOfTag + ", "
			}
		}
	}
	updateConditionString = strings.TrimSuffix(updateConditionString, ", ")
	return updateConditionString
}
