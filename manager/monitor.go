package manager

import (
	"expvar"
	"reflect"

	"github.com/BeameryHQ/kubeaware/types"
)

// structTag will inspect variables that are contained
const structTag = "monitor"

// exportVariables should always recieve a pointer to a struct
// so that the underlying code can export any non complex values to expvar
func exportVariables(m interface{}) error {
	// Obtain the variable using the reflection library
	// so that we can export these variables ready for monitoring
	abstract := reflect.ValueOf(m)
	if abstract.Kind() == reflect.Ptr {
		// Need to access the variables within the struct
		abstract = abstract.Elem()
	}
	for i := 0; i < abstract.NumField(); i++ {
		switch abstract.Field(i).Kind() {
		case reflect.Slice, reflect.Ptr, reflect.Map:
			// Do Nothing
		case reflect.Struct:
			if !abstract.Field(i).IsValid() || !abstract.Field(i).Addr().CanInterface() {
				continue
			}
			if err := exportVariables(abstract.Field(i).Addr().Interface()); err != nil {
				return err
			}
		default:
			if tag, exist := abstract.Type().Field(i).Tag.Lookup(structTag); exist {
				var variable types.Polymorph
				variable.Set(reflect.Indirect(reflect.ValueOf(m)).FieldByName(abstract.Type().Field(i).Name))
				expvar.Publish(tag, variable)
			}
		}
	}
	return nil
}
