package manager

import (
	"expvar"
	"fmt"
	"reflect"

	"github.com/BeameryHQ/kubeaware/types"
)

// structTag will inspect variables that are contained
const structTag = "monitor"

func exportVariables(m interface{}) error {
	// Obtain the variable using the reflection library
	// so that we can export these variables ready for monitoring
	abstract := reflect.ValueOf(m)
	if abstract.Kind() == reflect.Ptr {
		// Need to access al the variables within the struct
		abstract = abstract.Elem()
	}
	for i := 0; i < abstract.NumField(); i++ {
		switch abstract.Field(i).Kind() {
		case reflect.Struct:
			// inner structs doesn't work yet as its all magical
			if err := exportVariables(abstract.Field(i).Addr()); err != nil {
				return err
			}
			continue
		default:
			if tag, exist := abstract.Type().Field(i).Tag.Lookup(structTag); exist {
				var variable types.Polymorph
				variable.Set(reflect.Indirect(reflect.ValueOf(m)).FieldByName(abstract.Type().Field(i).Name))
				fmt.Println("Publishing tag", tag)
				expvar.Publish(tag, variable)
			}
		}
	}
	return nil
}
