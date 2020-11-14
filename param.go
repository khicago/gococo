package gococo

type Param string

func (p Param) String() string {
	return string(p)
}

func (p Param) IsNil() bool {
	return string(p) == NilParam
}
