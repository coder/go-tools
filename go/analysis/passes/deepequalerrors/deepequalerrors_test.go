// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deepequalerrors_test

import (
	"testing"

	"go.coder.com/go-tools/go/analysis/analysistest"
	"go.coder.com/go-tools/go/analysis/passes/deepequalerrors"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, deepequalerrors.Analyzer, "a")
}
