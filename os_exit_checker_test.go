package os_exit_checker

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestOsExitChecker(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, OsExitChecker, "./...")
}
