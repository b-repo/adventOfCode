package _8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedFirstInput = SevenSegmentDisplay{
	Signal: []Output{33, 510510, 255255, 170170, 2805, 85085, 102102, 15015, 2730, 231},
	Digit:  []Output{510510, 15015, 255255, 2805},
}

func TestReadSignalsAndDigits(t *testing.T) {
	inputs, err := ReadSignalsAndDigits("testdata/08.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expectedFirstInput, inputs[0])
}

func TestCountDigitsWithUniqueSegmentsAmount(t *testing.T) {
	inputs, err := ReadSignalsAndDigits("testdata/08.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 26, CountDigitsWithUniqueSegmentsAmount(inputs))
}

func TestSevenSegmentDisplay_Decrypt(t *testing.T) {
	inputs, err := ReadSignalsAndDigits("testdata/08.log")
	if err != nil {
		t.Fail()
	}

	expected := []string{"8394", "9781", "1197", "9361", "4873", "8418", "4548", "1625", "8717", "4315"}

	for i := range inputs {
		inputs[i].Decrypt()
		assert.Equal(t, expected[i], inputs[i].Read())
	}
}

func TestSumDrecryptedDigits(t *testing.T) {
	inputs, err := ReadSignalsAndDigits("testdata/08.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 61229, SumDrecryptedDigits(inputs))
}
