package gococo

import (
	"regexp"
	"strings"
)

// e.g <% Command_Name: param1, param2, param3 %>
const (
	coReDoc   = "\\<\\%\\s*([^\\s\\,\\:\\%\\~]+)\\s*\\:?((?:[ \\t]*[^\\s\\,\\:\\%\\~]+[ \\t]*\\,?)*)\\%\\>"
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


