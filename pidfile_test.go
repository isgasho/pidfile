package pidfile

import (
	"testing"
)

func TestPidfile(t *testing.T) {
	p := New("/tmp/test.pid")
	if err := p.Create(); err != nil {
		t.Fatal(err.Error())
	}
	defer p.Remove()

	if pid := p.ProcessID(); pid == 0 {
		t.Fatal("pidfile create failed")
	}
}
