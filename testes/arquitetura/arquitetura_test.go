package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Nao funciona em arquitura amd64") // teste ira ser pulado
	}
	t.Fail()
}
