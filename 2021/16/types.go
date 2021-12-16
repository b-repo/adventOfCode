package _16

import (
	"errors"
	"strconv"
)

var bitMap = map[string]int64{
	"0000": 0,
	"0001": 1,
	"0010": 2,
	"0011": 3,
	"0100": 4,
	"0101": 5,
	"0110": 6,
	"0111": 7,
	"1000": 8,
	"1001": 9,
	"1010": 10,
	"1011": 11,
	"1100": 12,
	"1101": 13,
	"1110": 14,
	"1111": 15,
}

type Bits []int

type Packet struct {
	Version Bits
	TypeID  Bits
	Wrapped Bits
}

type TypeID int64

const (
	Sum     TypeID = 0
	Product TypeID = 1
	Min     TypeID = 2
	Max     TypeID = 3
	Greater TypeID = 5
	Lesser  TypeID = 6
	Equal   TypeID = 7
)

func (b Bits) String() string {
	s := ""
	for i := range b {
		s += strconv.Itoa(b[i])
	}

	for len(s) < 4 {
		s = "0" + s
	}

	return s
}

func (b Bits) Decode() (int64, int64, int, error) {
	vsum := int64(0)
	p := b.ParsePacket()
	vsum += bitMap[p.Version.String()]
	t := bitMap[p.TypeID.String()]

	if t == 4 {
		value, left, err := p.Wrapped.ParseLiteralValue()

		return vsum, value, len(left), err
	}

	v, r, left, err := p.Wrapped.ParseOperator(TypeID(t))
	if err != nil {
		return 0, 0, 0, err
	}

	vsum += v

	return vsum, r, left, nil
}

func (b Bits) ParseOperator(t TypeID) (int64, int64, int, error) {
	i := b[0]
	if i == 0 {
		l := b[1:16].String()

		vl, err := strconv.ParseInt(l, 2, 64)
		if err != nil {
			return 0, 0, 0, err
		}

		subset := b[16 : 16+vl]
		left := b[16+vl:]

		vsum := int64(0)
		value := int64(-1)
		for len(subset) > 0 {
			v, r, lenLeft, err := subset.Decode()
			if err != nil {
				return 0, 0, 0, err
			}

			subset = subset[len(subset)-lenLeft:]

			value = Operate(t, value, r)
			vsum += v
		}

		return vsum, value, len(left), nil
	}

	if i == 1 {
		l := b[1:12].String()

		vl, err := strconv.ParseInt(l, 2, 64)
		if err != nil {
			return 0, 0, 0, err
		}

		subset := b[12:]

		vsum := int64(0)
		value := int64(-1)

		for n := int64(0); n < vl; n++ {
			v, r, left, err := subset.Decode()
			if err != nil {
				return 0, 0, 0, err
			}

			vsum += v
			value = Operate(t, value, r)
			subset = subset[len(subset)-left:]
		}

		return vsum, value, len(subset), err
	}

	return 0, 0, 0, errors.New("unable to parse operator")
}

func (b Bits) ParsePacket() Packet {
	return Packet{
		Version: b[:3],
		TypeID:  b[3:6],
		Wrapped: b[6:],
	}
}

func (b Bits) ParseLiteralValue() (int64, Bits, error) {
	isLastGroup := b[0] == 0
	v := b[1:]
	s := v[:4].String()
	v = v[4:]

	for !isLastGroup {
		isLastGroup = v[0] == 0
		v = v[1:]
		s += v[:4].String()
		v = v[4:]
	}

	value, err := strconv.ParseInt(s, 2, 64)

	return value, v, err
}
