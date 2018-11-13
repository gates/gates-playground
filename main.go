package main

import (
	"context"
	"time"

	"github.com/go-shadow/moment"
	"github.com/gopherjs/gopherjs/js"
	"github.com/lujjjh/gates"
)

func _time(fc gates.FunctionCall) gates.Value {
	args := fc.Args()
	layout := "X"
	if len(args) > 0 {
		layout = args[0].ToString()
	}
	return gates.String(moment.New().Format(layout))
}

func registerGlobal(r *gates.Runtime) {
	r.Global().InitBuiltIns()
	r.Global().Set("time", gates.FunctionFunc(_time))
}

func main() {
	js.Global.Set("RunString", func(s string) *js.Object {
		r := gates.New()
		registerGlobal(r)
		program, err := gates.Compile(s)
		if err != nil {
			panic(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		v, err := r.RunProgram(ctx, program)
		if err != nil {
			panic(err)
		}
		return js.MakeWrapper(v)
	})
}
