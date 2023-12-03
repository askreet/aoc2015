package day3

// Houses is a contiguous linear collection of houses designed to avoid pointer jumping or oversized
// memory allocation. It's probabyl better just to make a huge 2d grid and run the algorithm.
type Houses []House

func (hs *Houses) At(id int16) *House {
	return &(*hs)[id]
}

// Add is O(n) (linear search) for all existing houses when attaching existing exits.
func (hs *Houses) Add(x, y int16) int16 {
	*hs = append(*hs, House{
		X:     x,
		Y:     y,
		North: -1,
		South: -1,
		East:  -1,
		West:  -1,
	})

	newId := int16(len(*hs) - 1)
	conns := 0
	for id, house := range *hs {
		id := int16(id)

		if house.X == x && house.Y == y-1 {
			(*hs)[newId].North = id
			(*hs)[id].South = newId
			conns++
		}

		if house.X == x && house.Y == y+1 {
			(*hs)[newId].South = id
			(*hs)[id].North = newId
			conns++
		}

		if house.X == x-1 && house.Y == y {
			(*hs)[newId].West = id
			(*hs)[id].East = newId
			conns++
		}

		if house.X == x+1 && house.Y == y {
			(*hs)[newId].East = id
			(*hs)[id].West = newId
			conns++
		}

		if conns == 4 {
			return newId
		}
	}

	return newId
}

type House struct {
	North, South, East, West int16
	X, Y                     int16
	//Visits                   int
}

// NewHouse creates a new house with no linked exits.
func NewHouse() House {
	return House{
		North: -1,
		South: -1,
		East:  -1,
		West:  -1,
	}
}
