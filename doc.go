// Package goee is a DSL and widget set to make using Gio easier and smoother.
//
// todo: write a nice top level thing for the go.dev/pkg page
package goee

import (
	_ "embed"
	_ "gioui.org/app"
)

//go:embed version
var Version st
