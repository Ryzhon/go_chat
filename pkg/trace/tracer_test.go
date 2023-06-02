package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("returning value is nil from new")
	} else {
		tracer.Trace("hello trace package")
		if buf.String() != "hello trace package\n" {
			t.Errorf("'%s'output wrong string", buf.String())
		}
	}
}
