package manager

import (
	"expvar"
	"reflect"

	"github.com/BeameryHQ/kubeaware/types"
	"github.com/MovieStoreGuy/artemis"
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
		case reflect.Struct, reflect.Slice, reflect.Ptr, reflect.Map:
			// TODO(Sean Marciniak): Fix nested structs so that they can be evaluated
			artemis.GetInstance().Log(artemis.DebugEntry("Currently an unsupported type: ", abstract.Field(i).Kind()))
		default:
			if tag, exist := abstract.Type().Field(i).Tag.Lookup(structTag); exist {
				var variable types.Polymorph
				variable.Set(reflect.Indirect(reflect.ValueOf(m)).FieldByName(abstract.Type().Field(i).Name))
				artemis.GetInstance().Log(artemis.DebugEntry("Now publishing: ", tag))
				expvar.Publish(tag, variable)
			}
		}
	}
	return nil
}
