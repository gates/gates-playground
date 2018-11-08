package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lujjjh/gates"
)

func main() {
	js.Global.Set("RunString", func(s string) *js.Object {
		r := gates.New()
		v, err := r.RunString(s)
		if err != nil {
			panic(err)
		}
		return js.MakeWrapper(v)
	})
}
