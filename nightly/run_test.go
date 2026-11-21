package nightly

import "testing"

func TestRun(t *testing.T) {
	if Run44896913530() != 44896913530 {
		t.Fatal("wrong run")
	}
}
