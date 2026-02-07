package golife

import (
	"strings"
	"unicode"
)

func ReadRle(width int, height int, rleString string) *Field {
	data := make([]int, width*height)
	dataIdx := 0
	cumulator := 0

	for _, ch := range rleString {
		if unicode.IsDigit(ch) {
			cumulator = cumulator*10 + int(ch-'0')
			continue
		}
		count := 1
		if cumulator > 0 {
			count = cumulator
			cumulator = 0
		}

		for i := 0; i < count; i++ {
			if dataIdx < len(data) {
				if ch == '#' {
					data[dataIdx] = 1
				} else {
					data[dataIdx] = 0
				}
				dataIdx++
			}
		}
	}

	return &Field{
		cols: width,
		rows: height,
		data: data,
	}
}

func ReadStrings(s string) *Field {
	rows := strings.Split(strings.TrimSpace(s), "\n")
	var cleanRows []string

	for _, r := range rows {
		trimmed := strings.TrimSpace(r)
		if trimmed != "" {
			cleanRows = append(cleanRows, trimmed)
		}
	}

	if len(cleanRows) == 0 {
		return &Field{
			cols: 0,
			rows: 0,
			data: make([]int, 0),
		}
	}

	height := len(cleanRows)
	width := len(cleanRows[0])
	data := make([]int, height*width)
	result := Field{
		cols: width,
		rows: height,
		data: data,
	}
	for i, row := range cleanRows {
		row = strings.TrimRight(row, "\r")
		for j, ch := range row {
			alive := false
			if ch == '#' {
				alive = true
			}
			result.setAt(i, j, alive)
		}
	}
	return &result
}
