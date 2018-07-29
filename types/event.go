package types

type Condition int

const (
	NoOp          Condition = iota
	Shutdown      Condition = iota
	ForceShutdown Condition = iota
	Restart       Condition = iota
)
