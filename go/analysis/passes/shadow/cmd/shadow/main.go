// The shadow command runs the shadow analyzer.
package main

import (
	"go.coder.com/go-tools/go/analysis/passes/shadow"
	"go.coder.com/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
