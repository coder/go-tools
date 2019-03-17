// The findcall command runs the findcall analyzer.
package main

import (
	"go.coder.com/go-tools/go/analysis/passes/findcall"
	"go.coder.com/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(findcall.Analyzer) }
