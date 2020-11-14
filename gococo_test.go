package gococo_test

import (
	"fmt"
	"github.com/khicago/gococo"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func matchOrError(a, b interface{}, prefix string, t *testing.T) {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc)
	funcName := path.Base(caller.Name())
	fileName, line := caller.FileLine(pc)
	if !reflect.DeepEqual(a, b) {
		t.Errorf(
			"%s, value should match, expect %#v, got %#v.\n%s => %s:%d",
			prefix, a, b,
			funcName, fileName, line,
		)
	}
}

func TestNewCoco(t *testing.T) {
	co := gococo.NewCoco([]string{"cmd", "param1", "param2"})
	matchOrError("cmd", co.GetCMD(), "cmd error", t)
	matchOrError(gococo.Param("param1"), co.GetParams()[0], "param1 error", t)
	matchOrError(gococo.Param("param2"), co.GetParams()[1], "param2 error", t)
}

func TestDefaultSerializer(t *testing.T) {
	co := gococo.NewCoco([]string{"cmd", "param1", "param2"})
	matchOrError("<% cmd: param1, param2 %>", co.String(), "serializer error", t)
}

func TestParam_IsNil(t *testing.T) {
	co := gococo.NewCoco([]string{"cmd", gococo.NilParam, "another_val"})
	matchOrError(true, co.GetParams()[0].IsNil(), "param.IsNil error", t)
	matchOrError(false, co.GetParams()[1].IsNil(), "param.IsNil error", t)
}

func TestParse(t *testing.T) {
	strs := []string{
		"<% Test %>",
		"<% Num 3 4 5 _ 8 %>",
		"<% Arguments _,arg1,arg2,arg3 %>",
		"<% AllName alice,bob,cat,dada %>",
	}
	validate := []string{
		"<% Test %>",
		"<% Num: 3, 4, 5, _, 8 %>",
		"<% Arguments: _, arg1, arg2, arg3 %>",
		"<% AllName: alice, bob, cat, dada %>",
	}
	ret, ok := gococo.Parse(fmt.Sprintf(`
	this is a document
	for %s
	has %s
	and %s
	and %s
	//`, strs[0], strs[1], strs[2], strs[3]))

	if !ok {
		t.Error("parse failed")
	}

	if len(ret) != len(strs) {
		t.Fatalf("count of result error, expect= %d, got= %d", len(strs), len(ret))
	}

	for i, v := range ret {
		matchOrError(validate[i], v.String(), fmt.Sprintf("parse error, ind= %d", i), t)
	}
}
