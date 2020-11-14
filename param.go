package gococo

type Param string

// NilParam determines should a input will be considered a nil value.
// It will be called in parse
var NilParam = "_"

func (p Param) String() string {
	return string(p)
}

func (p Param) IsNil() bool {
	return string(p) == NilParam
}