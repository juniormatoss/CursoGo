package arquitetura

import (
	"runtime"
	"testing"
)

func TesteDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciomna em arquiterura amd64")
	}
	t.Fail()

}
