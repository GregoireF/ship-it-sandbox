package nightly

import "testing"

func TestRun(t *testing.T) {
	if Run43413091467() != 43413091467 {
		t.Fatal("wrong run")
	}
}
