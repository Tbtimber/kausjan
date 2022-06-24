package anp

import "testing"

func testHellow(t *testing.T) {
	if Hellow() != 1 {
		t.Fatalf("Error")
	}
}
