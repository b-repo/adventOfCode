package _16

import (
	"2021/common/reader"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBits_Decode(t *testing.T) {
	s, err := reader.ReadStringFileAndSplitByLine("testdata/16.log")
	if err != nil {
		t.Fail()
	}

	expectedVersionSum := []int64{6, 9, 14, 16, 12, 23, 31}

	for i := 0; i < len(expectedVersionSum); i++ {
		h, err := hex.DecodeString(s[i])
		if err != nil {
			t.Fail()
		}

		v, _, _, err := ParseBits(h).Decode()
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, expectedVersionSum[i], v)
	}
}

func TestBits_Decode2(t *testing.T) {
	s, err := reader.ReadStringFileAndSplitByLine("testdata/16_alt.log")
	if err != nil {
		t.Fail()
	}

	expectedValues := []int64{3, 54, 7, 9, 1, 0, 0, 1}

	for i := 4; i < len(expectedValues); i++ {
		h, err := hex.DecodeString(s[i])
		if err != nil {
			t.Fail()
		}

		_, v, _, err := ParseBits(h).Decode()
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, expectedValues[i], v)
	}
}
