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
	strs := []interface{}{
		"<% Test %>",
		"<% Num 3 4 5 _ 8 %>",
		"<% Arguments _,arg1,arg2,arg3 %>",
		"<% AllName alice,bob,cat,dada %>",
		"<%YOUR_CMD:_,PARAM1,PARAM2,PARAM3%>",
		"<% YOUR_CMD _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR_CMD %>",
		"<% YOUR_CMD _ PARAM1 PARAM2 PARAM3 %>",
		"<%    YOUR_CMD    :  _  ,  PARAM1  ,  PARAM2  ,   PARAM3   %>",
		"<% YOUR---CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR.@#*&^(){}<>./\\CMD: _, PARAM.@#*&^(){}<>./\\ %>",
		"<% YOUR一あ♂★CMD: _, PARAM.1, 一二三, あいう, 【】,♂★%>",
	}
	validate := []string{
		"<% Test %>",
		"<% Num: 3, 4, 5, _, 8 %>",
		"<% Arguments: _, arg1, arg2, arg3 %>",
		"<% AllName: alice, bob, cat, dada %>",
		"<% YOUR_CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR_CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR_CMD %>",
		"<% YOUR_CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR_CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR---CMD: _, PARAM1, PARAM2, PARAM3 %>",
		"<% YOUR.@#*&^(){}<>./\\CMD: _, PARAM.@#*&^(){}<>./\\ %>",
		"<% YOUR一あ♂★CMD: _, PARAM.1, 一二三, あいう, 【】, ♂★ %>",
	}
	testStr := fmt.Sprintf(`this is a document for %s has %s and %s
	and %s and %s and %s and %s and %s and %s and %s and %s and %s and %s and strange symbols <> !@#$^&*()~
	and %s and %s and %s and %s and %s and %s and %s and %s and %s and %s
	//`, strs ...)
	ret, ok := gococo.Parse(testStr)

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

func TestParseError(t *testing.T) {
	strs := []interface{}{
		"<% YOUR_CMD: _, PARAM1, PARAM2, PARAM3, %>",
		"<% YOUR_CMD: _, PAR~AM1, PARAM2, PARAM3, %>",
		"<% YOUR~CMD: _, PARAM1, PARAM2, PARAM3, %>",
		"<% YOUR_CMD: _, PARAM1, PARAM2, ,PARAM3, %>",
		`<% YOUR_CMD: _, PARAM1, 
		PARAM2, ,PARAM3, %>`,
	}
	for i, v := range strs {
		testStr := fmt.Sprintf(`//this is a document for %s //`, v)
		_, ok := gococo.Parse(testStr)
		if ok {
			t.Errorf("the %d case= %s should failed", i, v)
		}
	}
}
