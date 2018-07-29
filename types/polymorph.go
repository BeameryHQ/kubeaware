package types

import "fmt"

// Polymorph creates an object that satisfies the expvar.Var interface
// with more flexibility of what is stored inside it.
type Polymorph struct {
	val interface{}
}

// Set should only except values that are pointers to references
func (p *Polymorph) Set(v interface{}) {
	p.val = v
}

// String needs to satisfy the expvar.Var string's conversion
func (p Polymorph) String() string {
	return fmt.Sprintf("%v", p.val)
}
