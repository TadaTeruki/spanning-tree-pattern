package main

import (
	"math/rand"
	"syscall/js"
)

var PATTERN *Pattern

func setup(width, height, seed int) {
	rand.Seed(int64(seed))
	PATTERN = NewPattern(width, height)
}

func iterate() Iteration {
	return PATTERN.Iterate()
}

func main() {
	js.Global().Set("setupFunc", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		setup(args[0].Int(), args[1].Int(), args[2].Int())
		return nil
	}))

	js.Global().Set("iterateFunc", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		itr := iterate()
		data := make(map[string]interface{})
		data["x"] = itr.x
		data["y"] = itr.y
		data["depth"] = itr.depth
		data["hasValue"] = itr.hasValue
		return data
	}))

	done := make(chan struct{}, 0)
	<-done
}
