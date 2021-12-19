package _18

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	strings, err := reader.ReadStringFileAndSplitByLine("testdata/18.log")
	if err != nil {
		t.Fail()
	}

	r := strings[0]
	for i := 1; i < len(strings); i++ {
		r = Add(r, strings[i])
	}

	assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", r)

	n := TransformToTreeAndReturnMagnitude(r)

	assert.Equal(t, 4140, n)
}

func TestLargestMagnitude(t *testing.T) {
	strings, err := reader.ReadStringFileAndSplitByLine("testdata/18.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 3993, LargestMagnitude(strings))
}
