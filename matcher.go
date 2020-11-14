package gococo

import (
	"regexp"
	"strings"
)

type MatcherHandler func(strIn string) (results [][]string, find bool)

// e.g <% Command_Name: param1, param2, param3 %>
const coRegexp = "\\<\\%\\s*(\\w+)\\s*\\:?([ \\t]*\\w+[ \\t]*\\,?)*\\%\\>"

var exp, _ = regexp.Compile(coRegexp)

func defaultMatcher(doc string) ([][]string, bool) {
	if ind := strings.Index(doc, "<%"); ind < 0 {
		return nil, false
	} else {
		doc = doc[ind:]
	}

	mResult := exp.FindAllSubmatch([]byte(doc), -1)
	if len(mResult) <= 0 {
		return nil, false
	}

	rets := make([][]string, 0, len(mResult))
	for _, lineScan := range mResult {
		lenLineScan := len(lineScan)
		if lenLineScan <= 0 {
			continue
		}
		lineConv := make([]string, 0, lenLineScan - 1)
		for i := 1; i < lenLineScan; i ++ {
			lineConv = append(lineConv, string(lineScan[i]))
		}
		rets = append(rets, lineConv)
	}

	return rets, true
}


// Matcher is a function determines the rules of matching, and which can be set at implementation-side.
// It main function is receives a document and performs an matching procedure.
// After the matching regexp performed, each item of the result should be a slice of string.
// For each slice, the first value should indicates the command name, and followed by a list of param.
//
// The default matching can be found in testing files.
var Matcher MatcherHandler = defaultMatcher