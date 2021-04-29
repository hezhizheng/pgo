package pgo

import (
	"log"
	"testing"
)

func TestInArray(t *testing.T) {
	needle := "1"
	hystack := []string{"1", "2"}

	if !InArray(needle, hystack) {
		t.Error("test in_array fail")
	}
}

func TestArrayColumn(t *testing.T) {

	//map[string]map[string]interface{}

	input := map[string]map[string]interface{}{
		"0": {
			"x1": "1",
			"x2": "2",
		},
		"1": {
			"x1": "3",
			"x2": "4",
		},
	}

	columnKey := "x2"

	columns := ArrayColumn(input, columnKey)

	for _, v := range columns {
		if !(v == "4" || v == "2") {
			t.Error("test ArrayColumn fail", v)
		}
	}

	log.Println(columns, len(columns))
}
