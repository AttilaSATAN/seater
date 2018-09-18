package plane

import (
	"math"
	"strconv"
)

//Plane is the representation of the plane and its seats.
type Plane struct {
	ID           string
	ColNames     map[int]string
	RowNumber    int
	TotalSeats   int
	AislesAfter  []int
	SeatOrderMap []int
}

func (p Plane) SeatByIndex(i int) string {
	c := i % len(p.ColNames)
	r := math.Floor(float64(i / len(p.ColNames)))

	return strconv.Itoa(int(r)) + p.ColNames[c]
}

type PlanePrototype struct {
	ID             string
	SeatNumber     int
	ColumnNumber   int
	AisleAfterCols []int
}
