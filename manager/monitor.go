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
	// obtian the variable using the reflection library
	// so that we can export these variables ready for monitoring
	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Ptr {
		return errors.New("Unable monitor non pointer values")
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, exist := field.Tag.Lookup(structTag); exist {
			var variable types.Polymorph
			// Need to convert field to a pointer value so we can recieve updates once things
			// change within that variable
			// Need to ensure that the tag is valid json
			variable.Set(field)
			expvar.Publish(tag, variable)
		}
	}
	return nil
}
