package day1

import (
	"bufio"
	"io"
)

type Solution struct{}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	v := 0
	for scanner.Scan() {
		b := scanner.Bytes()[0]
		switch b {
		case '(':
			v++

		case ')':
			v--

		default:
			panic("unexpected scan: " + string(b))
		}
	}

	return v
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	v := 0
	pos := 1
	for scanner.Scan() {
		b := scanner.Bytes()[0]
		switch b {
		case '(':
			v++

		case ')':
			v--

		default:
			panic("unexpected scan: " + string(b))
		}
		if v == -1 {
			return pos
		}
		pos++
	}

	panic("Santa never reached the basement!")
}
