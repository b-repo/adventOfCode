package _21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlayerFromPosition(t *testing.T) {
	p1, p2, err := NewPlayerFromPosition("testdata/21.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, &Player{Position: 4}, p1)
	assert.Equal(t, &Player{Position: 8}, p2)
}

func TestPlay(t *testing.T) {
	p1, p2, err := NewPlayerFromPosition("testdata/21.log")
	if err != nil {
		t.Fail()
	}

	w, l, turns := Play(p1, p2, NewDeterministicDie(), 1000)

	assert.Equal(t, p1, w)
	assert.Equal(t, p2, l)
	assert.Equal(t, 993, turns*3)
	assert.Equal(t, 1000, p1.Score)
	assert.Equal(t, 745, p2.Score)
}

func TestPlayWithMultiverse(t *testing.T) {
	p1, p2, err := NewPlayerFromPosition("testdata/21.log")
	if err != nil {
		t.Fail()
	}

	w1, w2 := PlayWithMultiverse(p1, p2)

	assert.Equal(t, 444356092776315, w1)
	assert.Equal(t, 341960390180808, w2)
}
