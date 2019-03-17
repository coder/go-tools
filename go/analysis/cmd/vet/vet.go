// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The vet command is a static checker for Go programs. It has pluggable
// analyzers defined using the go.coder.com/go-tools/go/analysis API, and
// using the go.coder.com/go-tools/go/packages API to load packages in any
// build system.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"go.coder.com/go-tools/go/analysis/multichecker"

	// analysis plug-ins
	"go.coder.com/go-tools/go/analysis/passes/asmdecl"
	"go.coder.com/go-tools/go/analysis/passes/assign"
	"go.coder.com/go-tools/go/analysis/passes/atomic"
	"go.coder.com/go-tools/go/analysis/passes/atomicalign"
	"go.coder.com/go-tools/go/analysis/passes/bools"
	"go.coder.com/go-tools/go/analysis/passes/buildtag"
	"go.coder.com/go-tools/go/analysis/passes/cgocall"
	"go.coder.com/go-tools/go/analysis/passes/composite"
	"go.coder.com/go-tools/go/analysis/passes/copylock"
	"go.coder.com/go-tools/go/analysis/passes/httpresponse"
	"go.coder.com/go-tools/go/analysis/passes/loopclosure"
	"go.coder.com/go-tools/go/analysis/passes/lostcancel"
	"go.coder.com/go-tools/go/analysis/passes/nilfunc"
	"go.coder.com/go-tools/go/analysis/passes/nilness"
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
	// This suite of analyzers is applied to all code
	// in GOROOT by GOROOT/src/cmd/vet/all. When adding
	// a new analyzer, update the whitelist used by vet/all,
	// or change its vet command to disable the new analyzer.
	multichecker.Main(
		// the traditional vet suite:
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		atomicalign.Analyzer,
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

		// for debugging:
		// findcall.Analyzer,
		// pkgfact.Analyzer,

		// uses SSA:
		nilness.Analyzer,
	)
}
