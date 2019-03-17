// The lostcancel command applies the go.coder.com/go-tools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"go.coder.com/go-tools/go/analysis/passes/lostcancel"
	"go.coder.com/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
