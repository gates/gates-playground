package main

import (
	"github.com/go-shadow/moment"
	"github.com/gopherjs/gopherjs/js"
	"github.com/lujjjh/gates"
)

func time(fc gates.FunctionCall) gates.Value {
	args := fc.Args()
	layout := "X"
	if len(args) > 0 {
		layout = args[0].ToString()
	}
	return gates.String(moment.New().Format(layout))
}

func registerGlobal(r *gates.Runtime) {
	r.Global().InitBuiltIns()
	r.Global().Set("time", gates.FunctionFunc(time))
}

func main() {
	js.Global.Set("RunString", func(s string) *js.Object {
		r := gates.New()
		registerGlobal(r)
		v, err := r.RunString(s)
		if err != nil {
			panic(err)
		}
		return js.MakeWrapper(v)
	})
}
