package main

import (
	"log"
	"syscall/js"

	"github.com/vugu/vugu"
)

func (c *Root) ExecFromGo(e vugu.DOMEvent) {
	log.Printf("In ExecFromGo...\n")

	//j := js.Global().Call("jQuery", "#jquery-test-label")
	j := js.Global().Get("jQuery").New("#jquery-test-label")

	grey := "#cccccc"
	//j.SetCss("color", grey)

	colorMap := make(map[string]interface{})
	colorMap["color"] = grey
	j.Call("css", js.ValueOf(colorMap))

	t := j.Call("text")
	log.Printf("(Before update, from WASM) text = %v\n", t)
	j.Call("text", "Updated (from WASM)")

	t = j.Call("text")
	log.Printf("(After update, from WASM) text = %v\n", t)
}

func (c *Root) InvokeJS(e vugu.DOMEvent) {
	log.Printf("In InvokeJS...\n")

	h := js.Global().Get("helperFunc")
	h.Invoke()
}
