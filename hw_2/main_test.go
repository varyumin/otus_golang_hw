package main

import (
	"testing"
)

type Data struct {
	Req  string
	Resp string
}

func TestUnpack(t *testing.T) {
	data := []Data{
		{Req: "a4bc2d5e", Resp: "aaaabccddddde"},
		{Req: "abcd", Resp: "abcd"},
		{Req: "", Resp: ""},
	}
	for _, v := range data {
		result, err := Unpack(v.Resp)
		if err != nil {
			t.Errorf("Test FAILED: %s", err)
		}
		if result != v.Resp {
			t.Errorf("Invalid response %s != %s ", result, v.Resp)
		} else {
			t.Logf("\"%s\" == \"%s\"", result, v.Resp)
		}
	}
}
