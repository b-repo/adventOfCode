package _10

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFindIllegalCharacter(t *testing.T) {
	r, err := reader.Read2DStringSliceFile("testdata/10.log")
	if err != nil {
		t.Fail()
	}

	expectedIllegalChar := []string{"}", ")", "]", ")", ">"}
	expectedMissingChars := [][]string{
		{"}", "}", "]", "]", ")", "}", ")", "]"},
		{")", "}", ">", "]", "}", ")"},
		{"}", "}", ">", "}", ">", ")", ")", ")", ")"},
		{"]", "]", "}", "}", "]", "}", "]", "}", ">"},
		{"]", ")", "}", ">"},
	}
	j := 0
	k := 0
	for i := range r {
		c, p := FindErrorInLine(r[i], 0, nil)
		if c != "" {
			assert.Equal(t, expectedIllegalChar[j], c, "invalid response in line "+strconv.Itoa(i))
			j++
		} else {
			m := make([]string, 0)
			for p != nil {
				m = append(m, p.Opposite())
				p = p.Previous
			}
			assert.Equal(t, expectedMissingChars[k], m)
			k++
		}
	}
}

func TestCalculateSyntaxCheckerScore(t *testing.T) {
	assert.Equal(t, 26397, CalculateSyntaxCheckerScore([]string{"}", ")", "]", ")", ">"}))
}

func TestCalculateAutoCompleteScore(t *testing.T) {
	assert.Equal(t, 288957, CalculateAutoCompleteScore([][]string{
		{"}", "}", "]", "]", ")", "}", ")", "]"},
		{")", "}", ">", "]", "}", ")"},
		{"}", "}", ">", "}", ">", ")", ")", ")", ")"},
		{"]", "]", "}", "}", "]", "}", "]", "}", ">"},
		{"]", ")", "}", ">"},
	}))
}
