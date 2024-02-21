package main

import (
	"testing"
)


func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expect []string
    }{
        {
            input: "hello world",
            expect: []string{
            "hello",
            "world",
            },
        },
        {
            input: "HeLlo World",
            expect: []string{
            "hello",
            "world",
            },
        },
    }

    for _, val := range cases {
        res := cleanInput(val.input)
        if len(res) != len(val.expect) {
            t.Errorf("lengths does not match")
            return
        }
        for i,_ := range val.expect {
        if val.expect[i] != res[i]{
            t.Errorf("wrong result expected %v and got %v", val.expect[i], res[i])
            return
        }
    }

    }


}
