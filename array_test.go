package pgo

import (
	"github.com/stretchr/testify/assert"
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

// go test -v ./ -test.run TestArrayPush
func TestArrayPush(t *testing.T) {
	ary := &[]interface{}{
		"Winnie",
		"the",
		"Pooh",
	}

	ele := "Winnie-the-Pooh"

	push := ArrayPush(ary, ele)

	pushAry := *ary

	assert.Equal(t, 4, push)
	assert.Equal(t, pushAry[3], ele)
}

// go test -v ./ -test.run TestArrayPop
func TestArrayPop(t *testing.T) {
	ary := &[]interface{}{
		"Winnie",
		"the",
		"Pooh",
	}

	assert.True(t, InArray("Pooh",*ary))

	ArrayPop(ary)

	assert.Equal(t, 2, len(*ary))
	assert.True(t, !InArray("Pooh",*ary))
}

func TestArrayUnshift(t *testing.T) {
	ary := &[]interface{}{
		"Winnie",
		"the",
		"Pooh",
	}

	ele := "Winnie-the-Pooh"

	push := ArrayUnshift(ary, ele)

	pushAry := *ary

	assert.Equal(t, 4, push)
	assert.Equal(t, pushAry[0], ele)
}

func TestArrayShift(t *testing.T) {
	ary := &[]interface{}{
		"Winnie",
		"the",
		"Pooh",
	}

	assert.True(t, InArray("Winnie",*ary))

	ArrayShift(ary)

	assert.Equal(t, 2, len(*ary))
	assert.True(t, !InArray("Winnie",*ary))
}
