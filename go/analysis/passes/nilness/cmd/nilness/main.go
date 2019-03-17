// The nilness command applies the go.coder.com/go-tools/go/analysis/passes/nilness
// analysis to the specified packages of Go source code.
package main

import (
	"go.coder.com/go-tools/go/analysis/passes/nilness"
	"go.coder.com/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nilness.Analyzer) }
