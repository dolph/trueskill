package mathextra

import "math"

// Float64AlmostEq returns true if the difference between two floats is less than or equal to epsilon (ε).
func Float64AlmostEq(a, b, ε float64) bool {
	return math.Abs(a-b) <= ε
}
