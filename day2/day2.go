package day2

import (
	"bufio"
	"fmt"
	"io"
)

type Solution struct{}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var l, w, h, sum int
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		_, err := fmt.Sscanf(line, "%dx%dx%d\n", &l, &w, &h)
		if err != nil {
			panic(err)
		}

		sum += (2 * l * w) + (2 * w * h) + (2 * h * l) + min(l*w, w*h, h*l)
	}

	return sum
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var l, w, h, sum int
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		_, err := fmt.Sscanf(line, "%dx%dx%d\n", &l, &w, &h)
		if err != nil {
			panic(err)
		}

		sum += RibbonNeed(l, w, h)
	}

	return sum
}

func RibbonNeed(l, w, h int) int {
	return min(2*(l+w), 2*(w+h), 2*(h+l)) + (l * w * h)
}
