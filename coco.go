package gococo

import (
	"strings"
)

type (
	coco struct {
		Cmd    string  `json:"cmd"`
		Params []Param `json:"params"`
	}
)

func (c *coco) String() string {
	return Serializer(c)
}

func (c *coco) GetCMD() string {
	return c.Cmd
}

func (c *coco) GetParams() []Param {
	return c.Params
}

var _ Coco = &coco{}

// NewCoco create a new coco instance by a slice of []byte
func NewCoco(scanResult []string) Coco {
	lenRet := len(scanResult)
	if lenRet == 0 {
		return nil
	}

	cmd := &coco{
		Cmd:    strings.TrimSpace(scanResult[0]),
		Params: make([]Param, 0, lenRet-1),
	}
	// trust the Matcher function
	for _, param := range scanResult[1:] {
		cmd.Params = append(cmd.Params, Param(param))
	}
	return cmd
}

func defaultSerializer(co Coco) string {
	builder := strings.Builder{}
	builder.WriteString("<% ")
	builder.WriteString(co.GetCMD())
	lenP := len(co.GetParams())
	if lenP > 0 {
		builder.WriteRune(':')
	}
	for i, v := range co.GetParams() {
		builder.WriteRune(' ')
		builder.WriteString(v.String())
		if i != lenP-1 {
			builder.WriteRune(',')
		}
	}
	builder.WriteString(" %>")
	return builder.String()
}
