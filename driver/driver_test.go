package driver

import "testing"

func TestNew(t *testing.T) {
	if _, err := New("unknown://url", TxnPerFile); err == nil {
		t.Error("no error although driver unknown")
	}
}
