// The unmarshal command runs the unmarshal analyzer.
package main

import (
	"go.coder.com/go-tools/go/analysis/passes/unmarshal"
	"go.coder.com/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unmarshal.Analyzer) }
