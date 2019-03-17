package shadow_test

import (
	"testing"

	"go.coder.com/go-tools/go/analysis/analysistest"
	"go.coder.com/go-tools/go/analysis/passes/shadow"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, shadow.Analyzer, "a")
}
