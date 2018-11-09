package main

import (
	"github.com/go-shadow/moment"
	"github.com/gopherjs/gopherjs/js"
	"github.com/lujjjh/gates"
)

var global = gates.Map{
	"time": gates.FunctionFunc(func(fc gates.FunctionCall) gates.Value {
		args := fc.Args()
		layout := "X"
		if len(args) > 0 {
			layout = args[0].ToString()
		}
		return gates.String(moment.New().Format(layout))
	}),
}

func main() {
	js.Global.Set("RunString", func(s string) *js.Object {
		r := gates.New()
		r.SetGlobal(global)
		v, err := r.RunString(s)
		if err != nil {
			panic(err)
		}
		return js.MakeWrapper(v)
	})
}
