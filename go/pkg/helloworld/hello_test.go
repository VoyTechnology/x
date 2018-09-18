package helloworld

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	v := "world"
	if res := Hello(v); !strings.Contains(res, v) {
		t.Errorf("expected %s in the output, had %s", v, res)
	}
}
