package lib

import (
	"reflect"
	"testing"
)

func TestReadRle(t *testing.T) {
	tests := []struct {
		name       string
		cols, rows int
		rle        string
		wantData   []int
	}{
		{
			name: "Single cells",
			cols: 2, rows: 1,
			rle:      ".#",
			wantData: []int{0, 1},
		},
		{
			name: "Multi-digits RLE",
			cols: 5, rows: 1,
			rle:      "3#2.",
			wantData: []int{1, 1, 1, 0, 0},
		},
		{
			name: "Large multiplier",
			cols: 10, rows: 1,
			rle:      "10#",
			wantData: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "Mixed chars",
			cols: 4, rows: 1,
			rle:      "2#1x1?",
			wantData: []int{1, 1, 0, 0},
		},
		{
			name: "Overflow protection",
			cols: 2, rows: 1,
			rle:      "5#",
			wantData: []int{1, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := ReadRle(test.cols, test.rows, test.rle)
			if !reflect.DeepEqual(f.data, test.wantData) {
				t.Errorf("ReadRle(%q) = %v, want %v", test.rle, f.data, test.wantData)
			}
		})
	}

}

func TestSaveRle(t *testing.T) {
	tests := []struct {
		name       string
		sourceRle  string
		cols, rows int
	}{
		{
			name:      "Simple RLE",
			sourceRle: "5#",
			cols:      5, rows: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := ReadRle(test.cols, test.rows, test.sourceRle)
			res := SaveRle(*f)
			if test.sourceRle != res {
				t.Errorf("SaveRle(%q) = %v, want %v", test.sourceRle, res, test.sourceRle)
			}
		})
	}
}
