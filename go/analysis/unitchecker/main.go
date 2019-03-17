// +build ignore

// This file provides an example command for static checkers
// conforming to the go.coder.com/go-tools/go/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//   $ go build -o unitchecker main.go
//   $ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"go.coder.com/go-tools/go/analysis/unitchecker"

	"go.coder.com/go-tools/go/analysis/passes/asmdecl"
	"go.coder.com/go-tools/go/analysis/passes/assign"
	"go.coder.com/go-tools/go/analysis/passes/atomic"
	"go.coder.com/go-tools/go/analysis/passes/bools"
	"go.coder.com/go-tools/go/analysis/passes/buildtag"
	"go.coder.com/go-tools/go/analysis/passes/cgocall"
	"go.coder.com/go-tools/go/analysis/passes/composite"
	"go.coder.com/go-tools/go/analysis/passes/copylock"
	"go.coder.com/go-tools/go/analysis/passes/httpresponse"
	"go.coder.com/go-tools/go/analysis/passes/loopclosure"
	"go.coder.com/go-tools/go/analysis/passes/lostcancel"
	"go.coder.com/go-tools/go/analysis/passes/nilfunc"
	"go.coder.com/go-tools/go/analysis/passes/printf"
	"go.coder.com/go-tools/go/analysis/passes/shift"
	"go.coder.com/go-tools/go/analysis/passes/stdmethods"
	"go.coder.com/go-tools/go/analysis/passes/structtag"
	"go.coder.com/go-tools/go/analysis/passes/tests"
	"go.coder.com/go-tools/go/analysis/passes/unmarshal"
	"go.coder.com/go-tools/go/analysis/passes/unreachable"
	"go.coder.com/go-tools/go/analysis/passes/unsafeptr"
	"go.coder.com/go-tools/go/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		httpresponse.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		stdmethods.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
