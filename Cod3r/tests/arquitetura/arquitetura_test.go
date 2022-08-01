package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Skipping test, AMD64 only")
	}

	t.Fail()
}
