package types_test

import (
	"fmt"
	"testing"

	"github.com/BeameryHQ/kubeaware/types"
)

func TestPolymorph(t *testing.T) {
	p := types.Polymorph{}
	various := []interface{}{10, 3.14, 1e3, "str", struct{}{}}
	for count, v := range various {
		p.Set(v)
		if p.String() != fmt.Sprintf("%v", v) {
			t.Fatalf("Expecting that item can be converted at %v", count)
		}
	}
}
