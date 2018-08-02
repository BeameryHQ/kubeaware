package manager

import (
	"errors"
	"expvar"
	"reflect"

	"github.com/BeameryHQ/kubeaware/types"
)

// structTag will inspect variables that are contained
const structTag = "monitor"

func exportVariables(m types.Module) error {
	// Obtain the variable using the reflection library
	// so that we can export these variables ready for monitoring
	mirrortype := reflect.TypeOf(m)
	if mirrortype.Kind() != reflect.Ptr && mirrortype.Elem().Kind() != reflect.Struct {
		return errors.New("Unable monitor non pointer values")
	}
	mirrortype = mirrortype.Elem()
	for i := 0; i < mirrortype.NumField(); i++ {
		field := mirrortype.Field(i)
		if tag, exist := field.Tag.Lookup(structTag); exist {
			// May need to be smarter with embedded structs
			var variable types.Polymorph
			r := reflect.ValueOf(m)
			t := reflect.Indirect(r).FieldByName(field.Name)
			variable.Set(t)
			expvar.Publish(tag, variable)
		}
	}
	return nil
}
