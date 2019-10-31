package logger

import "testing"

func TestNew(t *testing.T) {
	if _, err := New("console"); err != nil {
		t.Errorf("invalid logger %+v", err)
	}
	if _, err := New("file"); err == nil {
		t.Error("should fail")
	}
}