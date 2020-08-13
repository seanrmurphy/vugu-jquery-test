// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vugu/vugu/devutil"
)

func readIndexFile() {
	indexFile := "./index.html"
	dat, err := ioutil.ReadFile(indexFile)
	if err != nil {
		panic(err)
	}

	devutil.DefaultIndex = devutil.StaticContent(string(dat))

	devutil.DefaultAutoReloadIndex = devutil.DefaultIndex.Replace(
		"<!-- scripts -->",
		"<script src=\"http://localhost:8324/auto-reload.js\"></script>\n<!-- scripts -->")
	//log.Printf("DefaultAutoReloadIndex = %v\n", devutil.DefaultAutoReloadIndex)
}

func main() {

	readIndexFile()

	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	wc := devutil.NewWasmCompiler().SetDir(".")
	mux := devutil.NewMux()

	mux.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
		`<!-- styles -->`,
		`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">`))
	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	mux.Default(devutil.NewFileServer().SetDir("."))

	log.Fatal(http.ListenAndServe(l, mux))
}
