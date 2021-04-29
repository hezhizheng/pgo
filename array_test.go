package pgo

import "testing"

func TestInArray(t *testing.T) {
	needle := "1"
	hystack := []string{"1","2"}

	if !InArray(needle,hystack) {
		t.Error("test in_array fail")
	}
}