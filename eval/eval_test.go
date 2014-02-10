package eval

import (
	"strconv"
	"syscall"
	"testing"
)

func strsEqual(s1 []string, s2 []string) bool {
	if len(s1) == len(s2) {
		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
			return true
		}
	}
	return false
}

func TestNewEvaluator(t *testing.T) {
	ev := NewEvaluator()
	pid := strconv.Itoa(syscall.Getpid())
	if ev.globals["pid"].String(ev) != pid {
		t.Errorf(`ev.globals["pid"] = %v, want %v`, ev.globals["pid"], pid)
	}
}
