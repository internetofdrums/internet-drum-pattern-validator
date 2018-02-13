package main

import (
	"bytes"
	"reflect"
	"testing"
)

var validDrumPatternBytes = []byte{
	0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x40, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var validDrumPattern = DrumPattern{
	[]Instrument{
		{[]Note{{0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x40, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x40, 0x40}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x40, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x40, 0x40}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x7F, 0x7F}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
		{[]Note{{0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}, {0x00, 0x00}}},
	},
}

var validDrumPatternString = `(7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) 
(00,00)(00,00)(40,7F)(00,00) (00,00)(00,00)(40,40)(00,00) (00,00)(00,00)(40,7F)(00,00) (00,00)(00,00)(40,40)(00,00) 
(00,00)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
`

func TestDecode(t *testing.T) {
	tests := []struct {
		in                string
		expectedOut       []byte
		shouldReturnError bool
	}{
		{
			"",
			[]byte{},
			false,
		},
		{
			"VGVzdA==",
			[]byte{0x54, 0x65, 0x73, 0x74},
			false,
		},
		{
			"VGVzdA==!",
			nil,
			true,
		},
	}

	for _, test := range tests {
		actualOut, err := Decode(test.in)
		hasReturnedError := err != nil

		if bytes.Compare(actualOut, test.expectedOut) != 0 {
			t.Errorf("Decode(%#v): expected <%#v>, actual <%#v>", test.in, test.expectedOut, actualOut)
		}

		if hasReturnedError != test.shouldReturnError {
			t.Errorf("Decode(%#v): expected error <%#v>, actual <%#v>", test.in, test.shouldReturnError, hasReturnedError)
		}
	}
}

func TestValidatePattern(t *testing.T) {
	tests := []struct {
		in                []byte
		shouldReturnError bool
	}{
		{
			nil,
			true,
		},
		{
			[]byte{},
			true,
		},
		{
			[]byte{0x00},
			true,
		},
		{
			[]byte{
				0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			true,
		},
		{
			validDrumPatternBytes,
			false,
		},
	}

	for _, test := range tests {
		err := ValidatePattern(test.in)
		hasReturnedError := err != nil

		if hasReturnedError != test.shouldReturnError {
			t.Errorf("ValidatePattern(%#v): expected error <%#v>, actual <%#v>", test.in, test.shouldReturnError, hasReturnedError)
		}
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		in          []byte
		expectedOut DrumPattern
	}{
		{
			validDrumPatternBytes,
			validDrumPattern,
		},
	}

	for _, test := range tests {
		actualOut := Convert(test.in)

		if !reflect.DeepEqual(actualOut, test.expectedOut) {
			t.Errorf("Convert(%#v): expected <%#v>, actual <%#v>", test.in, test.expectedOut, actualOut)
		}
	}
}

func TestGetFormattedPattern(t *testing.T) {
	tests := []struct {
		in          DrumPattern
		expectedOut string
	}{
		{
			validDrumPattern,
			validDrumPatternString,
		},
	}

	for _, test := range tests {
		actualOut := GetFormattedPattern(test.in)

		if actualOut != test.expectedOut {
			t.Errorf("GetFormattedPattern(%#v): expected <%#v>, actual <%#v>", nil, test.expectedOut, actualOut)
		}
	}
}
