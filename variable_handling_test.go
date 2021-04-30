package pgo

import (
	"log"
	"testing"
)

func TestBlank(t *testing.T) {
	empty := Blank(000)

	if empty {
		t.Error("test Empty fail")
	}

	log.Println(empty)
}

func TestIsNumeric(t *testing.T) {

	isNumeric := IsNumeric("000")

	if !isNumeric {
		t.Error("test IsNumeric fail")
	}

	log.Println(isNumeric)
}