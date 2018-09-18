package plane

import (
	"errors"
	"math"
	"sort"
)

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func New(pp PlanePrototype) (Plane, error) {
	p := Plane{}

	if pp.SeatNumber < 1 {
		return p, errors.New("planes should have seats")
	}

	if pp.ColumnNumber > 26 {
		return p, errors.New("can not be more than 26 columns of seats in a plane. Yeah! Thats a rule. Just because I'm lazy")
	}

	sort.Ints(pp.AisleAfterCols)

	u := make([]int, 0, len(pp.AisleAfterCols))
	m := make(map[int]bool)
	//uniqueness
	for _, val := range pp.AisleAfterCols {
		if val == 0 || val == pp.ColumnNumber {
			return p, errors.New("it would not be smart to put aisles between windows and seats")
		}
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		} else {
			return p, errors.New("aislesAfter does not contain unique column numbers")
		}
	}

	if len(pp.AisleAfterCols) > pp.ColumnNumber-1 {
		return p, errors.New("not enough columns for that aisles number")
	}

	p.ID = pp.ID
	colLetters := alphabet[:pp.ColumnNumber]
	p.ColNames = make(map[int]string)
	for n, c := range colLetters {
		p.ColNames[n+1] = c
	}
	p.TotalSeats = pp.SeatNumber
	p.RowNumber = int(math.Ceil(float64(p.TotalSeats / pp.ColumnNumber)))
	p.AislesAfter = pp.AisleAfterCols
	p.SeatOrderMap = createSeatOrderMap(p)
	return p, nil
}

func createSeatOrderMap(p Plane) []int {

	order := []int{}

	m := make(map[int]bool)

	// uniqueness of the order map
	for _, val := range p.AislesAfter {
		// while counting left to right seat before/after aisle
		if _, sba := m[val]; !sba {
			m[val] = true
			order = append(order, val)
		}

		if _, saa := m[val+1]; !saa {
			m[val+1] = true
			order = append(order, val+1)
		}
	}
	//windows
	order = append(order, 1, len(p.ColNames))
	m[1] = true
	m[len(p.ColNames)] = true
	//rest
	for i := 1; i <= len(p.ColNames); i++ {
		if _, added := m[i]; !added {
			order = append(order, i)
			m[i] = true
		}
	}
	return order
}
