package twocrystalballproblem

import "math"

// TwoCrystalball returns the highest floor from which a crystal ball can be dropped without breaking.
func TwoCrystalBall(breaks []bool) int {
	jmpAmount := int(math.Sqrt(float64((len(breaks)))))

	i := jmpAmount
	found := false

	for !found && i < len(breaks){
		if breaks[i] {
			found = true
		} else {
			i += jmpAmount
		}
	}

	i -= jmpAmount
	for j := 0; j < jmpAmount && i < len(breaks); j++{
		if breaks[i] {
			return i
		}
		i++
	}
	return -1
}