package core

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if v := Version(); !strings.HasPrefix(v, "mdi/1.") {
		t.Errorf("invalid version %v", v)
	}
}