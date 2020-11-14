package gococo

import (
	"regexp"
	"strings"
)

type MatcherHandler func(strIn string) (results [][]string, find bool)

// e.g <% Command_Name: param1, param2, param3 %>
const (
	coReDoc   = "\\<\\%\\s*(\\w+)\\s*\\:?((?:[ \\t]*\\w+[ \\t]*\\,?)*)\\%\\>"
	coReParam = "[ \\t]*[\\, ][ \\t]*"
)

var (
	expDoc, _   = regexp.Compile(coReDoc)
	expParam, _ = regexp.Compile(coReParam)
)

func defaultMatcher(doc string) ([][]string, bool) {
	if ind := strings.Index(doc, "<%"); ind < 0 {
		return nil, false
	} else {
		doc = doc[ind:]
	}

	mResults := expDoc.FindAllStringSubmatch(doc, -1)
	if len(mResults) <= 0 {
		return nil, false
	}

	rets := make([][]string, 0, len(mResults))
	for _, lineScan := range mResults {
		cmd := lineScan[1]
		params := expParam.Split(lineScan[2], -1)
		lenParams := len(params)
		if lenParams <= 0 {
			continue
		}
		lineConv := make([]string, 1, lenParams+1)
		lineConv[0] = cmd
		for i := 0; i < lenParams; i++ {
			param := strings.TrimSpace(params[i])
			lenParam := len(param)
			if lenParam <= 0 { // empty
				continue
			}
			lineConv = append(lineConv, param)
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
//
// About params parse:
// 1. Empty param are not allowed, the null value should be indicate by the NilPat.
// 2. CmdName and Params should be trim in the match procedure. (FYI, strings.TrimSpace can be used in such condition)
var Matcher MatcherHandler = defaultMatcher

