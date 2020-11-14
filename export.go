package gococo

import "fmt"

type (
	SerializeHandler func(co Coco) string
	MatcherHandler   func(strIn string) (results [][]string, find bool)

	Coco interface {
		fmt.Stringer
		GetCMD() string
		GetParams() []Param
	}
)

// Matcher is a function determines the rules of matching, and which can be set at implementation-side.
// It main function is receives a document and performs an matching procedure.
// After the matching regexp performed, each item of the result should be a slice of string.
// For each slice, the first value should indicates the command name, and followed by a list of param.
//
// The default matching can be found in testing files.
//
// About params parse:
// 1. Empty param are not allowed, the null value should be indicate by the NilPat.
// 2. CmdName and Params should be trim in the match procedure. (FYI, strings.TrimSpace can be used in such condition)
var Matcher MatcherHandler = defaultMatcher

var Serializer SerializeHandler = defaultSerializer

func Parse(doc string) ([]Coco, bool) {
	results, ok := Matcher(doc)
	if !ok {
		return nil, false
	}
	cmds := make([]Coco, 0, len(results))
	for _, ret := range results {
		cmds = append(cmds, NewCoco(ret))
	}
	return cmds, true
}
