package day3

import (
	"bufio"
	"io"
)

type Solution struct{}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)
	var houses Houses
	houses = append(houses, NewHouse())

	santa := Agent{X: 0, Y: 0, LocationID: 0}

	for scanner.Scan() {
		b := scanner.Bytes()
		santa.Move(b[0], &houses)
	}

	return len(houses)
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)
	var houses Houses
	houses = append(houses, NewHouse())

	agents := []Agent{
		{X: 0, Y: 0, LocationID: 0}, // Santa
		{X: 0, Y: 0, LocationID: 0}, // Robo-Santa
	}

	agentId := 0
	for scanner.Scan() {
		b := scanner.Bytes()

		agents[agentId].Move(b[0], &houses)

		agentId++
		if agentId == len(agents) {
			agentId = 0
		}
	}

	return len(houses)
}

type Agent struct {
	LocationID int16
	X, Y       int16
}

func (a *Agent) Move(b byte, houses *Houses) {
	switch b {
	case '^':
		a.Y--
		if houses.At(a.LocationID).North != -1 {
			a.LocationID = houses.At(a.LocationID).North
		} else {
			a.LocationID = houses.Add(a.X, a.Y)
		}
	case 'v':
		a.Y++
		if houses.At(a.LocationID).South != -1 {
			a.LocationID = houses.At(a.LocationID).South
		} else {
			a.LocationID = houses.Add(a.X, a.Y)
		}
	case '>':
		a.X++
		if houses.At(a.LocationID).East != -1 {
			a.LocationID = houses.At(a.LocationID).East
		} else {
			a.LocationID = houses.Add(a.X, a.Y)
		}
	case '<':
		a.X--
		if houses.At(a.LocationID).West != -1 {
			a.LocationID = houses.At(a.LocationID).West
		} else {
			a.LocationID = houses.Add(a.X, a.Y)
		}
	default:
		panic("unexpected input: " + string(b))
	}
}
