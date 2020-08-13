// setup.go

package main

import (
	"github.com/vugu/vugu"
)

// OVERALL APPLICATION WIRING IN vuguSetup
func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	// CREATE THE ROOT COMPONENT
	root := &Root{}
	buildEnv.WireComponent(root) // WIRE IT

	return root
}
